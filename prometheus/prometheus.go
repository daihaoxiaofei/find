package prometheus

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 尚未完善
func Run() {
	// 创建一个没有任何 label 标签的 gauge 指标
	temp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "home_temperature_celsius",
		Help: "The current temperature in degrees Celsius.",
	})

	// 在默认的注册表中注册该指标
	prometheus.MustRegister(temp)

	// 设置 gauge 的值为 随机数
	temp.Set(rand.Float64())

	// 对外提供/metrics接口
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(`prometheus 启动出错`)
	}
}
