package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/user"
	Redis "douyin/pkg/redis"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis/v8"
)

type CommentUserInfo struct {
	UserId        int64  `json:"user_id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
type CommentRedisInfo struct {
	CommentId  int64           `json:"comment_id"`
	User       CommentUserInfo `json:"user"`
	Content    string          `json:"content"`
	CreateDate string          `json:"create_data"`
}

// 对象的序列化和反序列化方法
func (u *CommentUserInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
func (u *CommentUserInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}
func (c *CommentRedisInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}
func (c *CommentRedisInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

// 结构体转换
func ToDbUser(u CommentUserInfo) db.User {
	var dbUser db.User
	dbUser.ID = uint(u.UserId)
	dbUser.Name = u.Name
	dbUser.FollowerCount = int(u.FollowerCount)
	dbUser.FollowingCount = int(u.FollowCount)
	return dbUser
}
func ToRedisUser(u db.User) CommentUserInfo {
	var userInfo CommentUserInfo
	userInfo.UserId = int64(u.ID)
	userInfo.Name = u.Name
	userInfo.FollowCount = int64(u.FollowingCount)
	userInfo.FollowerCount = int64(u.FollowerCount)
	userInfo.IsFollow = false
	return userInfo
}
func ToRedisComment(u CommentUserInfo, c db.Comment) CommentRedisInfo {
	var commentInfo CommentRedisInfo
	commentInfo.User = u
	commentInfo.CommentId = int64(c.ID)
	commentInfo.Content = c.Content
	commentInfo.CreateDate = c.CreatTime.String()
	return commentInfo
}
func ToDbComment(c CommentRedisInfo, vid int64) (db.Comment, error) {
	var dbComment db.Comment
	dbComment.ID = uint(c.CommentId)
	dbComment.Content = c.Content
	dbComment.VideoId = int(vid)
	dbComment.UserId = int(c.User.UserId)
	t, err := time.ParseInLocation("2006-01-02 15:04:05", c.CreateDate, time.Local)
	if err != nil {
		klog.Fatalf("时间字符串转换类型失败")
		return dbComment, err
	}
	dbComment.CreatTime = t
	return dbComment, nil
}

func getIdCacheName(idType string) string {
	return "id_" + idType
}

// 使用Lua脚本可以获取唯一的自增ID
func GetOneId(ctx context.Context, RedisDb *redis.Client, idType string) (int64, error) {
	key := getIdCacheName(idType)
	luaId := redis.NewScript(`
		local id_key = KEYS[1]
		local current = redis.call('get',id_key)
		if current == false then
			redis.call('set',id_key,1)
			return '1'
		end
		--redis.log(redis.LOG_NOTICE,' current:'..current..':')
		local result =  tonumber(current)+1
		--redis.log(redis.LOG_NOTICE,' result:'..result..':')
		redis.call('set',id_key,result)
		return tostring(result)
	`)
	n, err := luaId.Run(ctx, RedisDb, []string{key}, 2).Result()
	if err != nil {
		return -1, err
	} else {
		var ret string = n.(string)
		retint, err := strconv.ParseInt(ret, 10, 64)
		if err == nil {
			return retint, err
		} else {
			return -1, err
		}
	}
}

// 包装redis数据到RPC数据结构
func RedisPackComment(ctx context.Context, m CommentRedisInfo) (*comment.Comment, error) {
	comment_user := &user.User{
		Id:            m.User.UserId,
		Name:          m.User.Name,
		FollowCount:   &m.User.FollowCount,
		FollowerCount: &m.User.FollowerCount,
		IsFollow:      m.User.IsFollow,
	}
	comment := &comment.Comment{
		Id:         m.CommentId,
		User:       comment_user,
		Content:    m.Content,
		CreateDate: m.CreateDate,
	}
	return comment, nil
}

func RedisPackComments(ctx context.Context, ms []CommentRedisInfo) ([]*comment.Comment, error) {
	comments := make([]*comment.Comment, 0)
	for _, m := range ms {
		n, err := RedisPackComment(ctx, m)
		if err == nil {
			comments = append(comments, n)
		}
	}
	return comments, nil
}

// 添加评论列表到redis里面
func AddRedisCommentList(ctx context.Context, vid int64, ms []*comment.Comment) error {
	comment_map := make(map[string]interface{})
	for _, m := range ms {
		if m != nil {
			// 构造缓存的数据结构
			userRedisInfo := CommentUserInfo{
				UserId:        m.User.Id,
				Name:          m.User.Name,
				FollowCount:   *m.User.FollowCount,
				FollowerCount: *m.User.FollowerCount,
				IsFollow:      m.User.IsFollow,
			}
			commentRedisInfo := CommentRedisInfo{
				CommentId:  m.Id,
				User:       userRedisInfo,
				Content:    m.Content,
				CreateDate: m.CreateDate,
			}
			cid_string := strconv.Itoa(int(m.Id))
			cid_binary, err := commentRedisInfo.MarshalBinary()
			if err != nil {
				klog.Fatalf("缓存写入redis时数据序列化错误")
				return err
			}
			comment_map[cid_string] = cid_binary
		}
	}
	// 调用redis的API一次性写入
	vid_string := strconv.Itoa(int(vid))
	err := db.CommentRedis.HMSet(ctx, vid_string, comment_map).Err()
	if err != nil {
		klog.Fatalf("评论列表写入redis失败")
		return err
	}
	return nil
}

// 获取单个用户的信息
func GetUserFromRedis(ctx context.Context, user_id int64) (*db.User, error) {
	// 查询redis缓存是否有用户数据
	uids := make([]uint, 0)
	uids = append(uids, uint(user_id))
	user := Redis.GetUsersFromRedis(ctx, uids)
	if user == nil {
		// 从MySQL中获取消息
		user, err := db.GetUserById(user_id)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return user[0], nil
}
