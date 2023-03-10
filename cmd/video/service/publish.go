package service

import (
	"context"
	"strconv"
	"time"

	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/oss"
	"douyin/cmd/video/pack"
	"douyin/kitex_gen/video"
	"github.com/cloudwego/kitex/pkg/klog"
)

var SuccessMsg = "Success"

// PublishAction implements the VideoSrvImpl interface.
func PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	videoTable := &db.Video{}
	playUrl, coverUrl, err := oss.UploadVideo(&req.Data)
	if err != nil {
		klog.CtxErrorf(ctx, "UploadVideo Err:%v", err.Error())
		return nil, err
	}
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
		StatusMsg:  &SuccessMsg,
	}
	return
}

// PublishList implements the VideoSrvImpl interface.
func PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	publishList, err := db.GetPublishListByAuthorId(req.UserId)
	if err != nil {
		return nil, err
	}

	videoList, err := pack.Videos(ctx, publishList)
	if err != nil {
		return nil, err
	}
	resp = &video.DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  &SuccessMsg,
		VideoList:  videoList,
	}
	return
}
