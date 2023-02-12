package heco

import (
	"find/pkg/config"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"testing"
)

func TestGetBalance(t *testing.T) {
	c := NewClient(config.Source.HECO.URL[0], config.Source.HECO.InitPNum)

	var balanceHexa hexutil.Big
	err := c.cli.Call(&balanceHexa, "eth_getBalance", `0x53572674504e32349dd8535833a9c4679536eddd`, "latest")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(`ok`, (*big.Int)(&balanceHexa))
}
