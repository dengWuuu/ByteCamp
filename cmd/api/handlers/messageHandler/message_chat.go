package messageHandler

import (
	"context"
	"douyin/cmd/api/handlers"
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/pack"
	"douyin/kitex_gen/message"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"google.golang.org/protobuf/runtime/protoimpl"
	"time"
)

func MessageChat(ctx context.Context, c *app.RequestContext) {

	// 1、绑定http参数
	var param handlers.MessageChatParam
	err := c.Bind(&param)
	if err != nil {
		hlog.Infof("参数绑定失败")
		panic(err)
	}

	// 2、入参校验
	if param.ToUserId < 0 {
		handlers.SendResponse(c, pack.BuildGetUserResp(errno.ErrBind))
	}
	// 3、调用rpc
	resp, err := rpc.MessageChat(ctx, &message.DouyinMessageChatRequest{
		Token:    param.Token,
		ToUserId: param.ToUserId,
	})
	list := resp.MessageList
	messageList := packMessageList(list)
	if err != nil {
		hlog.Fatalf("调用Message rpc 调用失败")
		return
	}
	c.JSON(consts.StatusOK, utils.H{
		"status_code":  resp.StatusCode, // 状态码，0-成功，其他值-失败
		"status_msg":   resp.StatusMsg,  // 返回状态描述
		"message_list": messageList,     // 聊天记录
	})

}

func packMessageList(list []*message.Message) []*Message {
	var packList = make([]*Message, len(list))

	for i := range list {
		m := list[i]
		parse, err := time.Parse("2006-01-02 15:04:05", *m.CreateTime)
		if err != nil {
			hlog.Fatal("获取聊天记录时 时间转换错误")
			panic(err)
		}
		packList[i] = &Message{
			Id:         m.Id,
			Content:    m.Content,
			CreateTime: parse.Unix(),
		}
	}
	return packList
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                        // 消息id
	Content       string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                               // 消息内容
	CreateTime    int64  `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"` // 消息创建时间
}
