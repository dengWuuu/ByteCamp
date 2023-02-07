package service

import (
	"context"
	"strconv"

	"douyin/cmd/comment/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"

	"github.com/cloudwego/kitex/pkg/klog"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {
	// 判断redis缓存里面是否存在video对象
	video_id := req.VideoId
	vid_string := strconv.Itoa(int(video_id))
	vid_cnt, err := db.CommentRedis.Exists(s.ctx, vid_string).Result()
	if err != nil {
		klog.Fatalf("redis查找video对象出错")
		panic(err)
	}
	if vid_cnt > 1 {
		klog.Fatalf("video对象不唯一")
		panic(err)
	}
	if vid_cnt == 1 {
		// 存在video对象
		comment_data, err := db.CommentRedis.HGetAll(s.ctx, vid_string).Result()
		if err != nil {
			panic(err)
		}
		commentRedisData := make([]CommentRedisInfo, 0)
		for _, comment_binary := range comment_data {
			var commentRedisInfo CommentRedisInfo
			commentRedisInfo.UnmarshalBinary([]byte(comment_binary))
			commentRedisData = append(commentRedisData, commentRedisInfo)
		}
		// 打包成rpc通信使用的数据结构体
		comments, err := RedisPackComments(s.ctx, commentRedisData)
		if err != nil {
			return nil, err
		}
		return comments, err
	}
	// 直接从数据库读取数据
	res, err := db.GetCommentByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	comments, err := pack.Comments(s.ctx, res)
	if err != nil {
		return nil, err
	}
	// 将数据存储到redis里面
	err = AddRedisCommentList(s.ctx, video_id, comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
