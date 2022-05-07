package okex

import (
	"encoding/json"
	"errors"
	"find/glog"
	"find/source/helpfunc"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"net/url"
	"time"
)

type Okex struct {
	Client *rpc.Client
}

func (in *Okex) GetBalance(address string) (balance *big.Int, ok bool, err error) {
RE:
	balance, err = helpfunc.GetBalance(in.Client, address)
	if err != nil {
		er, ok := err.(rpc.HTTPError)
		// 请求过于频繁 需要休眠
		if ok && (er.StatusCode == 405 || er.StatusCode == 0) {
			// todo: 控制所有协程一起睡眠的模式没有实现  目前只能凑合这样
			time.Sleep(time.Second * 60) // 睡一分钟?
			goto RE
		}
		jer, ok := err.(*json.SyntaxError)
		// 请求过于频繁 需要休眠
		if ok && jer.Error() == `invalid character '<' looking for beginning of value` {
			// todo: 控制所有协程一起睡眠的模式没有实现  目前只能凑合这样
			time.Sleep(time.Second * 60) // 睡一分钟?
			goto RE
		}
		urlErr, ok := err.(*url.Error)
		// tool.SmartPrint(urlErr.Err.Error())
		// fmt.Println(urlErr.Err.Error())
		// tool.SmartPrint(urlErr.Err)
		// os.Exit(0)
		// 请求过于频繁 需要休眠
		if ok {
			if urlErr.Err.Error() == `dial tcp: lookup exchainrpc.okex.org: no such host` {
				return nil, false, errors.New("okex 网络连接异常: " + err.Error())
			}
			fmt.Println(`urlErr.Err.Error()`, urlErr.Err.Error())
		}
		glog.ErrorExist(`okex: 获取余额出现未知错误: ` + err.Error())
		time.Sleep(time.Second * 60) // 任何错误都睡一分钟继续?
		goto RE
		// fmt.Println(`reflect.TypeOf(err)`,reflect.TypeOf(err))
		// return nil, false, errors.New("okex获取余额错误 e.Call : " + err.Error())
	}

	if balance.Cmp(big.NewInt(0)) != 1 { // 1e16 暂时先放宽要求 有余额就录入
		return balance, false, nil
	}

	return balance, true, nil
}
