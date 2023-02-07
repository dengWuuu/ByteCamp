package service

import (
	"context"
	"time"

	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/video"
)

// GetUserFeed implements the VideoSrvImpl interface.
func GetUserFeed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	dbVideos, err := db.GetVideosByLastTime(time.Unix(req.GetLatestTime(), 0))
	if err != nil {
		return nil, err
	}
	videoList, err := pack.Videos(ctx, dbVideos)

	nextTime := time.Now()
	for _, v := range dbVideos {
		if nextTime.After(v.UploadTime) {
			nextTime = v.UploadTime
		}
	}
	t := nextTime.Unix()
	resp = &video.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  &SuccessMsg,
		VideoList:  videoList,
		NextTime:   &t,
	}
	return
}

// GetVideoById implements the VideoSrvImpl interface.
func GetVideoById(ctx context.Context, req *video.VideoIdRequest) (resp *video.Video, err error) {
	return
}
