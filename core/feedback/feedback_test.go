package feedback

import (
	"context"
	"find/core/keypair"
	"find/pkg/config"
	"find/pkg/glog"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"net/url"
	"testing"
)

func Test_struct(t *testing.T) {
	client, err := ethclient.Dial(config.Source.ThirdParty.URL[0])
	if err != nil {
		panic(`client err: ` + err.Error())
	}
	ad := keypair.CreateKey()
	ctx := context.Background()
	workFun := func() bool {
		addr := <-ad
		// Get the balance of an account
		account := common.HexToAddress(addr.Address)
		balance, err := client.BalanceAt(ctx, account, nil)
		if err != nil {
			rpcHTTPError, ok := err.(rpc.HTTPError)
			// 请求过于频繁 需要休眠
			if ok && rpcHTTPError.StatusCode == 429 {
				return false
			}

			urlError, ok := err.(*url.Error)
			if ok {
				// fmt.Printf("eT2: %#v\n", urlError)
				fmt.Printf("eT2.Err: %#v\n", urlError.Err)
				// fmt.Printf("eT2.Err: %#v\n", urlError.Err)
				return false
			}

			fmt.Printf("其他错误 T: %#v\n", err)

			glog.Error(`client.BalanceAt`, zap.Error(err))

			return false
		}
		if balance.Cmp(big.NewInt(0)) == 1 { // 1e16 暂时先放宽要求 有余额就录入
			fmt.Println("有余额:",
				addr.Address, addr.Private, balance, decimal.NewFromBigInt(balance, -18).Truncate(5)) // 25893180161173005034
		} else {
			// fmt.Println("检测", addr.Private) // 25893180161173005034
		}
		return true
	}
	fb := NewFeedback(workFun, `test`, 30, 0.1)

	fb.Run()
}
