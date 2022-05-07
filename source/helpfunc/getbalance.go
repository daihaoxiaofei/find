package helpfunc

import (
	"errors"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)
// var EMap sync.Map

func GetBalance(Client *rpc.Client, address string) (balance *big.Int, err error) {
	var balanceStr string
	err = Client.Call(&balanceStr, "eth_getBalance", address, "latest")
	if err != nil {
		return nil, err
	}
	if balanceStr == `Ox` {
		return balance, nil
	}
	balance, ok := new(big.Int).SetString(balanceStr[2:], 16)
	if !ok {
		return balance, errors.New("数据转换错误 new(big.Int).SetString: " + balanceStr[2:])
	}

	if balance.Cmp(big.NewInt(0)) != 1 { // 1e16 暂时先放宽要求 有余额就录入
		return balance, nil
	}

	return balance, nil
}
//
// func WriteErr(err error)(requestAgain bool){
//
//
// }
