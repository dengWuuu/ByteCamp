package prometheus

import (
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var KitexClientTracer = prometheus.NewClientTracer(":9099", "/kitexclient")

// * var KitexServerTracer = prometheus.NewServerTracer(":9098", "/kitexserver") useless
// * for some reason the server tracer will open new prometheus server instance
// * 每一个微服务都开启一个prometheus server服务来将metric数据暴露到http中
