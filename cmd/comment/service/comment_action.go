package service

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"douyin/cmd/comment/commentMq"
	"douyin/cmd/comment/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"

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
func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	// 根据请求创建新的评论
	if req.ActionType == 1 {
		commentModel := &db.Comment{
			VideoId: int(req.VideoId),
			UserId:  int(req.UserId),
			Content: *req.CommentText,
		}
		// 从redis获取唯一的自增ID
		// TODO 加入到配置文件里面
		id_type := "comment_generator"
		comment_id, err := GetOneId(s.ctx, db.CommentRedis, id_type)
		cid_string := strconv.Itoa(int(comment_id))
		if err != nil {
			klog.Fatalf("获取自增ID错误")
			return nil, err
		}
		// * 手动添加ID和创建时间
		commentModel.ID = uint(comment_id)
		commentModel.CreatedAt = time.Now()
		// 查找redis里面是否存在对应的video对象
		vid_string := strconv.Itoa(int(req.VideoId))
		vid_cnt, err := db.CommentRedis.Exists(s.ctx, vid_string).Result()
		// 下面都是严重的错误
		if err != nil {
			klog.Fatalf("查询redis视频对象出错")
			return nil, err
		}
		if vid_cnt > 1 {
			klog.Fatalf("video对象不唯一")
			return nil, err
		}
		if vid_cnt == 1 {
			// 存在video对象
			// 配置评论对象信息
			var userInfo CommentUserInfo
			var commentInfo CommentRedisInfo
			user_info, err := GetUserFromRedis(s.ctx, req.UserId)
			if err != nil {
				klog.Fatalf("创建评论过程中获取用户信息失败")
				return nil, err
			}
			// TODO 判断是否关注
			userInfo = ToRedisUser(*user_info)
			commentInfo = ToRedisComment(userInfo, *commentModel)
			comment_binary, err := commentInfo.MarshalBinary()
			if err != nil {
				klog.Fatalf("评论信息序列化失败")
				return nil, err
			}
			// 修改video的键值对
			err = db.CommentRedis.HSet(s.ctx, vid_string, cid_string, comment_binary).Err()
			if err != nil {
				klog.Fatalf("redis修改video对象失败")
				return nil, err
			}
			// 发送消息给MQ
			// * 这里要手动加上comment的ID和创建时间
			commentMessage := commentMq.CommentRmqMessage{
				UserId:     commentModel.UserId,
				VideoId:    commentModel.VideoId,
				Content:    commentModel.Content,
				CreateTime: commentModel.CreatedAt,
				ActionType: int(req.ActionType),
				CommentId:  int(comment_id),
			}
			msg, err := json.Marshal(commentMessage)
			if err != nil {
				klog.Fatalf("序列化添加评论请求参数失败")
				return nil, err
			}
			commentMq.CommentActionMqSend(msg)
			return pack.Comment(s.ctx, commentModel)
		}
		// 直接修改数据库
		commentModel.ID = uint(comment_id) // 添加ID
		err = db.CreateComment(s.ctx, commentModel)
		if err != nil {
			klog.Fatalf("评论数据直接写入数据库失败")
			return nil, err
		}
		return pack.Comment(s.ctx, commentModel)
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
			return nil, err
		}
		if vid_cnt > 1 {
			klog.Fatalf("video对象不唯一")
			return nil, err
		}
		if vid_cnt == 1 {
			// 存在vidoe对象
			// 首先获取对象
			comment_binary, err := db.CommentRedis.HGet(s.ctx, vid_string, cid_string).Result()
			if err != nil {
				klog.Fatalf("删除评论过程中获取评论信息出错")
				return nil, err
			}
			var commentRedisInfo CommentRedisInfo
			commentRedisInfo.UnmarshalBinary([]byte(comment_binary))
			// 删除对象
			err = db.CommentRedis.HDel(s.ctx, vid_string, cid_string).Err()
			if err != nil {
				klog.Fatalf("redis修改video对象失败")
				return nil, err
			}
			// 发送消息给MQ
			commentMessage := commentMq.CommentRmqMessage{
				UserId:     int(commentRedisInfo.User.UserId),
				VideoId:    int(req.VideoId),
				ActionType: int(req.ActionType),
				CommentId:  int(comment_id),
			}
			msg, err := json.Marshal(commentMessage)
			if err != nil {
				klog.Fatalf("序列化添加评论请求参数失败")
				return nil, err
			}
			commentMq.CommentActionMqSend([]byte(msg))
			dbComment, err := ToDbComment(commentRedisInfo, req.VideoId)
			if err != nil {
				klog.Fatalf("评论转换格式失败")
				return nil, err
			}
			return pack.Comment(s.ctx, &dbComment)
		}
		// 直接修改数据库
		// 首先还要获取对应的comment信息
		dbComment, err := db.GetCommentByCommentId(s.ctx, []int{comment_id})
		if dbComment == nil {
			return nil, err
		}
		err = db.DeleteCommentById(s.ctx, video_id, comment_id)
		if err != nil {
			klog.Fatalf("删除评论直接删除出现错误")
			return nil, err
		}
		return pack.Comment(s.ctx, dbComment[0])
	}
	// 参数不合法
	return nil, errno.ErrBind
}
