package main

import (
	"find/core/keypair"
	"find/core/source"
	"find/pkg/glog"
)

func main() {
	glog.Info("开始运行")
	// 监控
	// go prometheus.Run()

	// 产出地址
	ad := keypair.CreateKey()

	// 资源
	sourcesArr := source.GetSources()
	if len(sourcesArr) == 0 {
		panic(`没有配置可用资源`)
	}
	for _, so := range sourcesArr { // 客户端资源组
		go so.Find(ad)
	}
	select {}
}
