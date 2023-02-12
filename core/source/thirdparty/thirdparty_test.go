package thirdparty

import (
	"find/core/keypair"
	"find/pkg/config"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
)

func TestGetBalance(t *testing.T) {
	c := NewClient(config.Source.Bsc.URL[1], config.Source.ThirdParty.InitPNum)

	var balanceHexa hexutil.Big
	err := c.cli.Call(&balanceHexa, "eth_getBalance", `0x53572674504e32349dd8535833a9c4679536eddd`, "latest")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(`ok`, (*big.Int)(&balanceHexa))
}

func TestFind(t *testing.T) {
	c := NewClient(config.Source.Bsc.URL[1], config.Source.Bsc.InitPNum)

	// 产出地址
	ad := keypair.CreateKey()
	c.Find(ad)
}
