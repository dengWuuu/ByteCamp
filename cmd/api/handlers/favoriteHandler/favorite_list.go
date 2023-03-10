package favoriteHandler

import (
	"context"
	"strconv"

	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/pack"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// FavoriteList 传送http请求上下文到rpc客户端，并且获得客户端的响应
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var favoriteListParam handlers.FavoriteListParam
	token := c.Query("token")
	user_id := c.Query("user_id")
	// 检查参数
	uid, err := strconv.Atoi(user_id)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	// 封装请求到rpc客户端
	favoriteListParam.UserId = int64(uid)
	favoriteListParam.Token = token
	queryParam := &favorite.DouyinFavoriteListRequest{
		UserId: favoriteListParam.UserId,
		Token:  favoriteListParam.Token,
	}
	rpcResp, err := rpc.FavoriteList(ctx, queryParam)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	c.JSON(200, utils.H{
		"status_code": rpcResp.StatusCode,
		"status_msg":  rpcResp.StatusMsg,
		"video_list":  rpcResp.VideoList,
	})
}
