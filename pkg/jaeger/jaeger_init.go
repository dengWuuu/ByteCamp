/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-08 13:39:58
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-08 13:40:22
 * @FilePath: /ByteCamp/pkg/jaeger/jaeger_init.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package jaeger

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	internal_opentracing "github.com/kitex-contrib/tracer-opentracing"
)

// InitJaeger ...
func InitJaegerServer(service string) (server.Suite, io.Closer) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return internal_opentracing.NewDefaultServerSuite(), closer
}

// InitJaeger ...
func InitJaegerClient(service string) (client.Suite, io.Closer) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return internal_opentracing.NewDefaultClientSuite(), closer
}
