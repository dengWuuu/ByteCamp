package service

import (
	"context"
	"douyin/cmd/comment/commentMq"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type CommentActionService struct {
	ctx context.Context
}

// 创建一个评论服务
func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// 评论服务实现
func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) error {
	// 根据请求创建新的评论
	if req.ActionType == 1 {
		commentModel := &db.Comment{
			VideoId:   int(req.VideoId),
			UserId:    int(req.UserId),
			Content:   *req.CommentText,
			CreatTime: time.Now(),
		}
		// 从redis获取唯一的自增ID
		// TODO 加入到配置文件里面
		id_type := "comment_generator"
		comment_id, err := GetOneId(s.ctx, db.CommentRedis, id_type)
		cid_string := strconv.Itoa(int(comment_id))
		if err != nil {
			klog.Fatalf("获取自增ID错误")
			return err
		}
		// 查找redis里面是否存在对应的video对象
		vid_string := strconv.Itoa(int(req.VideoId))
		vid_cnt, err := db.CommentRedis.Exists(s.ctx, vid_string).Result()
		// 下面都是严重的错误
		if err != nil {
			klog.Fatalf("查询redis视频对象出错")
			panic(err)
		}
		if vid_cnt > 1 {
			klog.Fatalf("video对象不唯一")
			panic(err)
		}
		if vid_cnt == 1 {
			// 存在video对象
			// 配置评论对象信息
			var userInfo CommentUserInfo
			var commentInfo CommentRedisInfo
			user_info, err := GetUserFromRedis(s.ctx, req.UserId)
			if err != nil {
				klog.Fatalf("创建评论过程中获取用户信息失败")
				return err
			}
			// TODO 判断是否关注
			userInfo = ToRedisUser(*user_info)
			commentInfo = ToRedisComment(userInfo, *commentModel)
			comment_binary, err := commentInfo.MarshalBinary()
			if err != nil {
				klog.Fatalf("评论信息序列化失败")
				return err
			}
			// 修改video的键值对
			err = db.CommentRedis.HSet(s.ctx, vid_string, cid_string, comment_binary).Err()
			if err != nil {
				klog.Fatalf("redis修改video对象失败")
				return err
			}
			// 发送消息给MQ
			req.CommentId = &comment_id
			msg, err := json.Marshal(req)
			if err != nil {
				klog.Fatalf("序列化添加评论请求参数失败")
				return err
			}
			commentMq.CommentActionMqSend([]byte(msg))
			return nil
		}
		// 直接修改数据库
		commentModel.ID = uint(comment_id) // 添加ID
		return db.CreateComment(s.ctx, commentModel)
	}
	// 根据请求删除评论
	if req.ActionType == 2 {
		comment_id := int(*req.CommentId)
		video_id := int(req.VideoId)
		vid_string := strconv.Itoa(video_id)
		cid_string := strconv.Itoa(comment_id)
		// 查找redis中是否存在
		vid_cnt, err := db.CommentRedis.Exists(s.ctx, vid_string).Result()
		// 下面都是严重的错误
		if err != nil {
			klog.Fatalf("查询redis video对象出错")
			panic(err)
		}
		if vid_cnt > 1 {
			klog.Fatalf("video对象不唯一")
			panic(err)
		}
		if vid_cnt == 1 {
			// 存在vidoe对象
			err := db.CommentRedis.HDel(s.ctx, vid_string, cid_string).Err()
			if err != nil {
				klog.Fatalf("redis修改video对象失败")
				return err
			}
			// 发送消息给MQ
			msg, err := json.Marshal(req)
			if err != nil {
				klog.Fatalf("序列化添加评论请求参数失败")
				return err
			}
			commentMq.CommentActionMqSend([]byte(msg))
			return nil
		}
		// 直接修改数据库
		return db.DeleteCommentById(s.ctx, video_id, comment_id)
	}
	// 参数不合法
	return errno.ErrBind
}
