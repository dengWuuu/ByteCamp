/*
 * @Author: Wu
 * @Date: 2022-1-30 14:14:40
 * @Description: 使用 Hertz 作为 http 服务器将请求转发到 RPC 服务器中
 */

// 使用 Hertz 提供 API 服务将 HTTP 请求发送给 RPC 微服务端
package main

import (
	"crypto/tls"
	"douyin/cmd/api/handlers/messageHandler"
	"os"

	"douyin/cmd/api/handlers/commentHandler"
	"douyin/cmd/api/handlers/favoriteHandler"
	"douyin/cmd/api/handlers/relationHandler"
	"douyin/cmd/api/handlers/userHandler"
	"douyin/cmd/api/handlers/videoHandler"
	"douyin/cmd/api/rpc"
	"douyin/dal"
	"douyin/pkg/middleware"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

func Init() {
	rpc.InitRpc() // 初始化rpc客户端
}

func InitNacos() naming_client.INamingClient {
	// nacos
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("81.70.207.243", 8848),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
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

// InitHertz 初始化 Hertz
func InitHertz() *server.Hertz {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("apiConfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/config")
	errV := viper.ReadInConfig()
	if errV != nil {
		hlog.Fatal("启动http服务器时读取配置文件失败")
		return nil
	}

	// Nacos
	cli := InitNacos()
	r := nacos.NewNacosRegistry(cli)
	opts := []config.Option{
		server.WithHostPorts(viper.GetString("Server.address") + ":" + viper.GetString("Server.Port")),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "API",
			Addr:        utils.NewNetAddr("tcp", viper.GetString("Server.address")+":"+viper.GetString("Server.Port")),
			Weight:      10,
			Tags:        nil,
		}),
	}

	// TLS & Http2
	tlsEnable := viper.GetBool("Hertz.Tls.Enable")
	h2Enable := viper.GetBool("Hertz.Http2.Enable")
	tlsConfig := tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
	if tlsEnable {
		cert, err := tls.LoadX509KeyPair(viper.GetString("Hertz.Tls.CertFile"), viper.GetString("Hertz.Tls.KeyFile"))
		if err != nil {
			hlog.Error(err)
		}
		tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
		opts = append(opts, server.WithTLS(&tlsConfig))

		if alpn := viper.GetBool("Hertz.Tls.ALPN"); alpn {
			opts = append(opts, server.WithALPN(alpn))
		}
	} else if h2Enable {
		opts = append(opts, server.WithH2C(h2Enable))
	}

	// Hertz
	h := server.Default(opts...)

	// JWT中间键初始化
	middleware.InitJwt()
	err = middleware.JwtMiddleware.MiddlewareInit()
	if err != nil {
		hlog.Fatalf("Jwt初始化失败")
		return nil
	}
	return h
}

// registerGroup 注册 Router组
func registerGroup(h *server.Hertz) {
	douyin := h.Group("/douyin")

	user := douyin.Group("/user")
	{
		// user模块下无需权限认证的接口
		user.POST("/register/", userHandler.Register)
		user.POST("/login/", userHandler.Login)

		// user模块下需要认证权限的接口
		user.Use(middleware.JwtMiddlewareFunc())
		{
			user.GET("/", userHandler.GetUserById)
		}
	}

	// relation模块接口
	relation := douyin.Group("/relation")
	relation.Use(middleware.JwtMiddleware.MiddlewareFunc())
	{
		relation.POST("/action/", relationHandler.RelationAction)
		relation.GET("/follow/list/", relationHandler.FollowList)
		relation.GET("/follower/list/", relationHandler.FollowerList)
		relation.GET("/friend/list/", relationHandler.FriendList)
	}

	// comment模块http接口
	comment := douyin.Group("/comment")
	comment.Use(middleware.JwtMiddleware.MiddlewareFunc())
	{
		comment.POST("/action/", commentHandler.CommentAction)
		comment.GET("/list/", commentHandler.CommentList)
	}
	// favorite模块http接口
	favorite := douyin.Group("/favorite")
	favorite.Use(middleware.JwtMiddleware.MiddlewareFunc())
	{
		favorite.POST("/action/", favoriteHandler.FavoriteAction)
		favorite.GET("/list/", favoriteHandler.FavoriteList)
	}

	// video模块接口
	douyin.GET("/feed/", videoHandler.Feed)
	publish := douyin.Group("/publish")
	publish.Use(middleware.JwtMiddlewareFunc())
	{
		publish.POST("/action/", videoHandler.PublishAction)
		publish.GET("/list/", videoHandler.PublishList)
	}
	// message模块接口
	message := douyin.Group("/message")
	message.Use(middleware.JwtMiddlewareFunc())
	{
		message.GET("/chat/", messageHandler.MessageChat)
		message.POST("/action/", messageHandler.MessageAction)
	}
}

// 初始化 Hertz服务器和路由组（Router）
func main() {
	// 数据库初始化
	dal.Init()
	// 设置系统日志框架 使用zap
	logger := hertzzap.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetSystemLogger(logger)

	Init()
	h := InitHertz()
	h.Use(middleware.AccessLog())
	pprof.Register(h)
	registerGroup(h)
	h.Spin()
}
