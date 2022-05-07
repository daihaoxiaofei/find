package source

import (
	"find/config"
	"find/source/bsc"
	"find/source/eth"
	"find/source/heco"
	"find/source/okex"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)


type Source interface {
	GetBalance(address string) (balance *big.Int, ok bool, err error)
}

// 根据配置文件决定用哪些数据源
func GetSources() (Sources []Source) {
	// todo: 未解决用变量名作为属性的方案
	// for _,v:=range []string{`Eth`}{
	// 	for _, url := range config.Source.Urls.(v) {
	//
	// 		c, err := rpc.Dial(url)
	// 		if err != nil {
	// 			panic(fmt.Sprintf("连接%v,错误 rpc.Dial %v", url, err))
	// 		}
	// 		Sources = append(Sources, &eth.Eth{Client: c})
	// 	}
	// }

	for _, url := range config.Source.Urls.Eth {
		c, err := rpc.Dial(url)
		if err != nil {
			panic(fmt.Sprintf("连接%v,错误 rpc.Dial %v", url, err))
		}
		Sources = append(Sources, &eth.Eth{Client: c})
	}

	for _, url := range config.Source.Urls.Bsc {
		c, err := rpc.Dial(url)
		if err != nil {
			panic(fmt.Sprintf("连接%v,错误 rpc.Dial %v", url, err))
		}
		Sources = append(Sources, &bsc.Bsc{Client: c})
	}

	for _, url := range config.Source.Urls.HECO {
		c, err := rpc.Dial(url)
		if err != nil {
			panic(fmt.Sprintf("连接%v,错误 rpc.Dial %v", url, err))
		}
		Sources = append(Sources, &heco.Heco{Client: c})
	}

	for _, url := range config.Source.Urls.OKEX {
		c, err := rpc.Dial(url)
		if err != nil {
			panic(fmt.Sprintf("连接%v,错误 rpc.Dial %v", url, err))
		}
		Sources = append(Sources, &okex.Okex{Client: c})
	}
	return
}
