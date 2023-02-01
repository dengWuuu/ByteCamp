/*
 * @Author: Wu
 * @Date: 2022-1-30 14:14:40
 * @Description: 使用 Hertz 作为 http 服务器将请求转发到 RPC 服务器中
 */

// 使用 Hertz 提供 API 服务将 HTTP 请求发送给 RPC 微服务端
package main

import (
	"crypto/tls"
	relationhandler "douyin/cmd/api/handlers/relationHandler"
	"douyin/cmd/api/handlers/userHandler"
	"douyin/cmd/api/rpc"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/pprof"
	"github.com/spf13/viper"
)

func Init() {
	rpc.InitRpc() //初始化rpc客户端
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
	opts := []config.Option{server.WithHostPorts(viper.GetString("Server.address") + ":" + viper.GetString("Server.Port"))}

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
	return h
}

// registerGroup 注册 Router组
func registerGroup(h *server.Hertz) {
	douyin := h.Group("/douyin")

	user := douyin.Group("/user")
	user.POST("/register/", userHandler.Register)
	user.GET("/", userHandler.GetUserById)

	//relation模块接口
	relation := douyin.Group("/relation")
	{
		relation.POST("/action", relationhandler.RelationAction)
		relation.GET("/follow/list", relationhandler.FollowList)
	}
}

// 初始化 Hertz服务器和路由组（Router）
func main() {
	//设置系统日志框架 使用zap
	logger := hertzzap.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetSystemLogger(logger)

	Init()
	h := InitHertz()
	pprof.Register(h)
	registerGroup(h)
	h.Spin()
}
