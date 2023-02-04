package favoriteHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/pack"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"douyin/pkg/middleware"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// FavoriteAction 传递请求上下文到favorite服务的rpc客户端并获取对应的响应
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var favoriteActionParam handlers.FavoriteActionParam
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	user_id := middleware.GetUserIdFromJwtToken(ctx, c)

	// 检查转换是否出错
	vid, err := strconv.Atoi(video_id)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	act, err := strconv.Atoi(action_type)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}

	favoriteActionParam.UserId = int64(user_id)
	favoriteActionParam.Token = token
	favoriteActionParam.VideoId = int64(vid)
	favoriteActionParam.ActionType = int32(act)

	// 检查参数
	if favoriteActionParam.UserId <= 0 || favoriteActionParam.VideoId <= 0 {
		handlers.SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	if favoriteActionParam.ActionType != 1 && favoriteActionParam.ActionType != 2 {
		handlers.SendResponse(c, pack.BuildFavoriteActionResp(errno.ErrBind))
		return
	}
	// 封装请求传送到rpc客户端
	queryParam := &favorite.DouyinFavoriteActionRequest{
		UserId:     favoriteActionParam.UserId,
		Token:      favoriteActionParam.Token,
		VideoId:    favoriteActionParam.VideoId,
		ActionType: favoriteActionParam.ActionType,
	}
	rpcResp, err := rpc.FavoriteAction(ctx, queryParam)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, rpcResp)
}
