package db

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Video struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	UploadTime    time.Time
	Title         string
	Version       optimisticlock.Version
	FavoriteCount int64
	CommentCount  int64
}

const VideoCount = 10

// GetVideosByLastTime
// 依据一个时间，来获取这个时间之前的一些视频
func GetVideosByLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, VideoCount)
	result := DB.Where("upload_time<?", lastTime).Order("upload_time desc").Limit(VideoCount).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	return videos, nil
}

// GetPublishListByAuthorId 获取用户发布列表
func GetPublishListByAuthorId(authorId int64) ([]*Video, error) {
	var data []*Video
	//初始化db
	//Init()
	result := DB.Where(&Video{AuthorId: authorId}).Find(&data)
	//如果出现问题，返回对应到空，并且返回error
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

// Save 保存视频记录
func Save(video Video) error {
	result := DB.Save(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
