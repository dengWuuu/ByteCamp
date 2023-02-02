package pack

import (
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"errors"
)

// 构建返回的响应状态信息
func BuildFavoriteActionResp(err error) *favorite.DouyinFavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteActionResp(e)
	}
	// unknown error
	e = errno.ErrUnknown.WithMessage(err.Error())
	return favoriteActionResp(e)
}
func BuildFavoriteListResp(err error) *favorite.DouyinFavoriteListResponse {
	if err == nil {
		return favoriteListResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return favoriteListResp(e)
	}
	// unknown error
	e = errno.ErrUnknown.WithMessage(err.Error())
	return favoriteListResp(e)
}

// 自定义构建的响应
func favoriteActionResp(err errno.ErrNo) *favorite.DouyinFavoriteActionResponse {
	return &favorite.DouyinFavoriteActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
func favoriteListResp(err errno.ErrNo) *favorite.DouyinFavoriteListResponse {
	return &favorite.DouyinFavoriteListResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
