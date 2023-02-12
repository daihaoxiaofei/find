package okex

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"find/pkg/config"
)

func TestGetBalance(t *testing.T) {
	c := NewClient(config.Source.OKEX.URL[0], config.Source.OKEX.InitPNum)

	var balanceHexa hexutil.Big
	err := c.cli.Call(&balanceHexa, "eth_getBalance", `0x53572674504e32349dd8535833a9c4679536eddd`, "latest")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(`ok`, (*big.Int)(&balanceHexa))
}
