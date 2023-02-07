package userService

import (
	"context"
	"log"
	"testing"

	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/usersrv"
	"github.com/cloudwego/kitex/client"
)

func TestRpcGetUser(t *testing.T) {
	c, err := usersrv.NewClient("test", client.WithHostPorts("127.0.0.1:8081"))
	if err != nil {
		log.Fatal(err)
	}
	req := &user.DouyinUserRequest{
		UserId: 1,
	}

	resp, _ := c.GetUserById(context.Background(), req)
	log.Println(resp)
}
