package favoriteHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/pack"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// FavoriteAction 传递请求上下文到favorite服务的rpc客户端并获取对应的响应
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var favoriteActionParam handlers.FavoriteActionParam
	err := c.Bind(&favoriteActionParam)
	if err != nil {
		hlog.Fatal("序列化喜欢行动请求参数失败")
		panic(err)
	}
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
