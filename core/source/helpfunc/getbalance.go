package helpfunc

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"

	"find/core/feedback"
	"find/core/keypair"
	"find/pkg/glog"
	"find/pkg/notify"
)

func succeedFun(from string, pair keypair.Pair, balance *big.Int) {
	str := fmt.Sprintf("from: %s\t 私钥为: %s\t地址为: %s\t余额为:%v 约:%s",
		from, pair.Private, pair.Address, balance, decimal.NewFromBigInt(balance, -18).Truncate(5))

	notify.PushMsg(str)
	glog.Ok(str)
}

// Find 在对应的资源中查找
func Find(cli *rpc.Client, ad <-chan keypair.Pair, from string, initPNum int, checkErr func(err error)) {
	workFun := func() bool {
		pair := <-ad
		var balanceHex hexutil.Big
		err := cli.Call(&balanceHex, "eth_getBalance", pair.Address, "latest")
		if err != nil {
			checkErr(err)
			return false
		}

		balance := (*big.Int)(&balanceHex)
		// fmt.Println("已检查", c.from, pair, balance)
		// 过滤小额 1e16 暂时先放宽要求 有余额就录入
		if balance.Cmp(big.NewInt(0)) == 1 {
			succeedFun(from, pair, balance)
		}
		return true
	}

	fb := feedback.NewFeedback(workFun, from, initPNum, 0.1)
	fb.Run()
}
