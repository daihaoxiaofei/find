package heco

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"testing"
)

func TestGetBalance(t *testing.T) {
	url := `https://bsc-dataseed.binance.org/`
	c, err := rpc.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("连接%v,错误 rpc.Dial %v", url, err))
	}
	b := &Heco{Client: c}

	fmt.Println(b.GetBalance(`0xA547381bA6c2b80c382bFBB338A2121D00907c5A`))
}
