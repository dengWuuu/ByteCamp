/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 01:04:01
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-07 16:34:45
 * @FilePath: \ByteCamp\cmd\relation\client\main.go
 * @Description: 用于测试relation微服务
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package main

import (
	"context"
	"fmt"

	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/relation/relationsrv"

	"github.com/cloudwego/kitex/client"
)

func main() {
	c, _ := relationsrv.NewClient("relation", client.WithHostPorts("127.0.0.1:8082"))
	request := relation.DouyinRelationFollowListRequest{
		UserId: 23,
		Token:  "",
	}
	resp, err := c.RelationFollowList(context.Background(), &request)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("resp: %v", resp)
}
