package service

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/video"
	"strconv"
	"time"
)

// PublishAction implements the VideoSrvImpl interface.
func PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {

	videoTable := &db.Video{}
	playUrl, coverUrl, err := db.UploadVideo(&req.Data)
	videoTable.AuthorId, _ = strconv.ParseInt(req.Token, 10, 64)
	videoTable.Title = req.Title
	videoTable.PlayUrl = playUrl
	videoTable.CoverUrl = coverUrl
	videoTable.UploadTime = time.Now()
	err = db.Save(*videoTable)
	if err != nil {
		return nil, err
	}
	resp = &video.DouyinPublishActionResponse{
		StatusCode: 0,
	}
	return
}

// PublishList implements the VideoSrvImpl interface.
func PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	resp = &video.DouyinPublishListResponse{
		StatusCode: 0,
	}
	return
}
