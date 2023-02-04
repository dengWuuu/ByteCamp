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
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"os"
)

var (
	VideoRPCPSM    string
	FavoriteRPCPSM string
	CommentRPCPSM  string
	RelationRPCPSM string
	UserRPCPSM     string
)

func InitRpc() {
	ReadConfig()

	initUserRpc()
	initRelationRpc()
	initCommentRpc()
	initFavoriteRpc()
	initVideoRpc()
}

func ReadConfig() {
	path, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动rpc客户端时读取配置文件失败")
		return
	}

	UserRPCPSM = viper.GetString("rpc.user.psm")
	RelationRPCPSM = viper.GetString("rpc.relation.psm")
	CommentRPCPSM = viper.GetString("rpc.comment.psm")
	FavoriteRPCPSM = viper.GetString("rpc.favorite.psm")
	VideoRPCPSM = viper.GetString("rpc.video.psm")
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
