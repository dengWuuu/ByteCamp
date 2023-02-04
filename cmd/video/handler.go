package main

import (
	"context"
	"douyin/cmd/video/service"
	video "douyin/kitex_gen/video"
	"github.com/cloudwego/kitex/pkg/klog"
)

// VideoSrvImpl implements the last service interface defined in the IDL.
type VideoSrvImpl struct{}

// PublishAction implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	klog.CtxInfof(ctx, "PublishAction Req: %v", req)
	resp, err = service.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	klog.CtxInfof(ctx, "PublishAction Resp: %v", resp)
	return
}

// PublishList implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	klog.CtxInfof(ctx, "PublishList Req: %v", req)
	resp, err = service.PublishList(ctx, req)

	if err != nil {
		return nil, err
	}

	klog.CtxInfof(ctx, "PublishList Resp: %v", resp)
	return
}

// GetUserFeed implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) GetUserFeed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	klog.CtxInfof(ctx, "GetUserFeed Req: %v", req)
	resp, err = service.GetUserFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	klog.CtxInfof(ctx, "GetUserFeed Resp: %v", resp)
	return
}

// GetVideoById implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) GetVideoById(ctx context.Context, req *video.VideoIdRequest) (resp *video.Video, err error) {
	resp, err = service.GetVideoById(ctx, req)
	return
}
