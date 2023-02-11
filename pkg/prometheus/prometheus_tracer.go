package prometheus

import (
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var KitexClientTracer = prometheus.NewClientTracer(":9099", "/kitexclient")

// TODO: 自定义服务指标和Tracer接口
// TODO: 使用Sentinel实现限流
// * 每一个微服务都开启一个prometheus server服务来将metric数据暴露到http中
