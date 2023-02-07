package userService

import (
	"context"
	"fmt"
	"log"
	"path"
	"runtime"
	"testing"

	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/usersrv"

	"github.com/cloudwego/kitex/client"
)

func TestRpcRegistry(t *testing.T) {
	c, err := usersrv.NewClient("test", client.WithHostPorts("127.0.0.1:8081"))
	if err != nil {
		log.Fatal(err)
	}
	req := &user.DouyinUserRegisterRequest{
		Username: "",
		Password: "",
	}

	resp, _ := c.Register(context.Background(), req)
	log.Println(resp)
}

func TestDir(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	fmt.Println(root)
}
