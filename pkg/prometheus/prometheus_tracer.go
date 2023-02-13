package prometheus

import (
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var KitexClientTracer = prometheus.NewClientTracer(":9099", "/kitexclient")
