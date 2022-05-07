package main

import (
	"find/config"
	"find/glog"
	"find/keypair"
	"find/notify"
	"find/prometheus"
	"find/source"
	"fmt"
	"sync"
)

func main() {
	// 监控
	go prometheus.Run()

	glog.Info("开始运行")
	// 产出地址
	ad := keypair.CreateKey()

	// 查询余额
	Sources := source.GetSources()
	wg := sync.WaitGroup{}
	for k := range Sources { // 客户端资源组
		for n := 0; n < config.Run.NumLine; n++ { // 多线程 客户端副本数
			wg.Add(1)
			go func() {
				Find(Sources[k], ad)
				wg.Done()
			}()

		}
	}
	wg.Wait()
}

func Find(so source.Source, ad <-chan keypair.Pair) {
	for {
		addr := <-ad
		Balance, ok, err := so.GetBalance(addr.Address)
		if err != nil {
			glog.Error(`获取余额出现错误:`, err)
			break
		}
		// glog.Infof("私钥为: %v\t地址为: %v\t余额为: %v", addr.Private, addr.Pair, Balance)
		if ok == false {
			continue
		}
		str := fmt.Sprintf("私钥为: %v\t地址为: %v\t余额为:%v", addr.Private, addr.Address, Balance)
		notify.PushMsg(str)
		glog.Error(str) // 虽然不是错误 但比较重要
	}
}
