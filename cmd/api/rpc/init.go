/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 12:23:15
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 02:21:45
 * @FilePath: /ByteCamp/cmd/api/rpc/init.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */

package rpc

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitRpc() {
	initUserRpc()
	initRelationRpc()
	initCommentRpc()
	initFavoriteRpc()
	initVideoRpc()
}

func NacosInit() naming_client.INamingClient {
	// the nacos server config
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("81.70.207.243", 8848),
	}

	// the nacos client config
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
		// more ...
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	return cli
}
