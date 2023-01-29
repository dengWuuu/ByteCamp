package main

import (
	"context"
	video "douyin/kitex_gen/video"
)

// VideoSrvImpl implements the last service interface defined in the IDL.
type VideoSrvImpl struct{}

// PublishAction implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) PublishList(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserFeed implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) GetUserFeed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoById implements the VideoSrvImpl interface.
func (s *VideoSrvImpl) GetVideoById(ctx context.Context, req *video.VideoIdRequest) (resp *video.Video, err error) {
	// TODO: Your code here...
	return
}
