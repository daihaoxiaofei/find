package heco

import (
	"find/glog"
	"find/source/helpfunc"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"time"
)

type Heco struct {
	Client *rpc.Client
}

func (in *Heco) GetBalance(address string) (balance *big.Int, ok bool, err error) {
RE:
	balance, err = helpfunc.GetBalance(in.Client, address)
	if err != nil {
		er, ok := err.(rpc.HTTPError)
		// fmt.Println(er)
		// os.Exit(0)
		// 请求过于频繁 需要休眠
		if ok && (er.StatusCode == 429 || er.StatusCode == 503 || er.StatusCode == 500) {
			// todo: 控制所有协程一起睡眠的模式没有实现  目前只能凑合这样
			time.Sleep(time.Second * 60) // 睡一分钟?
			goto RE
		}
		glog.ErrorExist(`heco: 获取余额出现未知错误: `+ err.Error())
		time.Sleep(time.Second * 60) // 任何错误都睡一分钟继续?
		goto RE
		// return nil, false, errors.New("heco获取余额错误 e.Call : " + err.Error())
	}

	if balance.Cmp(big.NewInt(0)) != 1 { // 1e16 暂时先放宽要求 有余额就录入
		return balance, false, nil
	}

	return balance, true, nil
}
