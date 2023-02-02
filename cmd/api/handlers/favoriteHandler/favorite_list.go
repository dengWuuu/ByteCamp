package favoriteHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/favorite/pack"
	"douyin/kitex_gen/favorite"
	"douyin/pkg/errno"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// 传送http请求上下文到rpc客户端，并且获得客户端的响应
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var favoriteListParam handlers.FavoriteListParam
	body, err := c.Body()
	if err != nil {
		hlog.Fatalf("获取请求体失败")
		panic(err)
	}
	err = json.Unmarshal(body, &favoriteListParam)
	if err != nil {
		hlog.Fatal("序列化评论获取请求参数失败")
		panic(err)
	}
	// 检查参数
	if favoriteListParam.UserId <= 0 {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ErrBind))
		return
	}
	// 封装请求到rpc客户端
	queryParam := &favorite.DouyinFavoriteListRequest{
		UserId: favoriteListParam.UserId,
		Token:  favoriteListParam.Token,
	}
	rpcResp, err := rpc.FavoriteList(ctx, queryParam)
	if err != nil {
		handlers.SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}
	handlers.SendResponse(c, rpcResp)
}
