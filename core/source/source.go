package source

import (
	"find/core/keypair"
	"find/core/source/bsc"
	"find/core/source/heco"
	"find/core/source/okex"
	"find/core/source/thirdparty"
	"find/pkg/config"
)

type Source interface {
	Find(ad <-chan keypair.Pair)
}

// GetSources 根据配置文件决定用哪些数据源
func GetSources() (sourcesArr []Source) {
	// 各各eth相关资源间  目前只有对错误的处理不同而已
	for _, url := range config.Source.ThirdParty.URL {
		sourcesArr = append(sourcesArr, thirdparty.NewClient(url, config.Source.ThirdParty.InitPNum))
	}
	for _, url := range config.Source.Bsc.URL {
		sourcesArr = append(sourcesArr, bsc.NewClient(url, config.Source.Bsc.InitPNum))
	}
	for _, url := range config.Source.HECO.URL {
		sourcesArr = append(sourcesArr, heco.NewClient(url, config.Source.HECO.InitPNum))
	}
	for _, url := range config.Source.OKEX.URL {
		sourcesArr = append(sourcesArr, okex.NewClient(url, config.Source.OKEX.InitPNum))
	}

	return
}
