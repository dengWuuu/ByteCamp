package relationhandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/relation/pack"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func FollowList(ctx context.Context, c *app.RequestContext) {
	var param handlers.FollowListParam
	//1、绑定http参数
	body, err := c.Body()
	if err != nil {
		hlog.Fatalf("获取请求体失败")
		panic(err)
	}
	err = json.Unmarshal(body, &param)
	if err != nil {
		hlog.Fatal("序列化关注列表请求参数失败")
		panic(err)
	}
	//2、入参校验
	if param.UserId == 0 {
		handlers.SendResponse(c, pack.BuildRelationFollowingListResp(nil, errno.ErrBind))
		return
	}

	//3、调用rpc
	resp, err := rpc.FollowList(ctx, &relation.DouyinRelationFollowListRequest{
		UserId: param.UserId,
		Token:  param.Token,
	})
	if err != nil {
		handlers.SendResponse(c, pack.BuildRelationFollowingListResp(nil, errno.ErrBind))
		return
	}
	handlers.SendResponse(c, resp)
}
