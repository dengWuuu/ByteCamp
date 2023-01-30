/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 01:04:01
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-31 01:26:09
 * @FilePath: /ByteCamp/cmd/relation/client/main.go
 * @Description: 用于测试relation微服务
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"context"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationsrv"
	"fmt"

	"github.com/cloudwego/kitex/client"
)

func main() {
	c, _ := relationsrv.NewClient("relation", client.WithHostPorts("127.0.0.1:8888"))
	request := relation.DouyinRelationActionRequest{
		UserId:     1,
		ToUserId:   2,
		ActionType: 2,
	}
	resp, err := c.RelationAction(context.Background(), &request)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("resp: %v", resp)
}
