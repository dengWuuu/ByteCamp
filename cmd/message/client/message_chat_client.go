package client

import (
	"douyin/kitex_gen/user/usersrv"
	"github.com/cloudwego/kitex/client"
	"log"
	"testing"
)

func TestRpcMessageChat(t *testing.T) {
	c, err := usersrv.NewClient("test", client.WithHostPorts("127.0.0.1:8081"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c)
}
