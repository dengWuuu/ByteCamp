package service

import (
	"context"
	"douyin/cmd/favorite/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"douyin/kitex_gen/video"
	"github.com/cloudwego/kitex/pkg/klog"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService 创建一个获取点赞视频服务
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{ctx: ctx}
}

// FavoriteList get the favorite list of user
func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*video.Video, error) {
	// 判断缓存里面是否有对象
	userId := req.UserId
	//userString := strconv.Itoa(int(userId))

	// 分为两个步骤：获取video信息，再打包成rpc的数据结构信息
	favorites, err := db.GetFavoritesByUserId(userId)
	if err != nil {
		klog.Errorf("mysql 获取用户喜欢视频列表失败")
		panic(err)
	}
	// 用户喜欢的视频列表ID
	var vIds = make([]int64, len(favorites))
	for index, f := range favorites {
		vIds[index] = int64(f.VideoId)
	}

	// 需要从数据库中调取
	videoMysql, err := db.GetVideoByIds(vIds)
	if err != nil {
		klog.Errorf("mysql 获取视频信息出错")
		panic(err)
	}
	videoData, err := pack.Videos(s.ctx, userId, videoMysql)
	if err != nil {
		klog.Errorf("将数据库打包成rpc数据格式出错")
		panic(err)
	}

	return videoData, nil
}
