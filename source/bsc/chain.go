package bsc

import (
	"find/glog"
	"find/source/helpfunc"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"time"
)

type Bsc struct {
	Client *rpc.Client
}

func (in *Bsc) GetBalance(address string) (balance *big.Int, ok bool, err error) {
RE:
	balance, err = helpfunc.GetBalance(in.Client, address)
	if err != nil {
		er, ok := err.(rpc.HTTPError)
		// 请求过于频繁 需要休眠
		if ok && (er.StatusCode == 403 || er.StatusCode == 405) { // 403 Forbidden.. 405 Not Allowed
			// todo: 控制所有协程一起睡眠的模式没有实现  目前只能凑合这样
			time.Sleep(time.Second * 60) // 睡一分钟?
			goto RE
		}
		glog.ErrorExist(`bsc: 获取余额出现未知错误: `+ err.Error())
		time.Sleep(time.Second * 60) // 任何错误都睡一分钟继续?
		goto RE
		// return nil, false, errors.New("bsc  e.Call : " + err.Error())
	}

	if balance.Cmp(big.NewInt(0)) != 1 { // 1e16 暂时先放宽要求 有余额就录入
		return balance, false, nil
	}

	return balance, true, nil
}
