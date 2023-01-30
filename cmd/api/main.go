/*
 * @Author: Wu
 * @Date: 2022-1-30 14:14:40
 * @Description: 使用 Hertz 作为 http 服务器将请求转发到 RPC 服务器中
 */

// 使用 Hertz 提供 API 服务将 HTTP 请求发送给 RPC 微服务端
package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/pprof"
)

// Init 初始化 API 配置
func Init() {

}

func InitHertzCfg() {

}

// InitHertz 初始化 Hertz
func InitHertz() *server.Hertz {
	InitHertzCfg()
	opts := []config.Option{server.WithHostPorts("127.0.0.1:8088")}
	// Hertz
	h := server.Default(opts...)
	return h
}

// 注册 Router组
func registerGroup(h *server.Hertz) {

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
