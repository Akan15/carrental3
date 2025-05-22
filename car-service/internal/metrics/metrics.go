package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Количество HTTP-запросов",
		},
		[]string{"handler", "method"},
	)
)

func Init() {
	prometheus.MustRegister(RequestCount)
}
