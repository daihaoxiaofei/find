package eth

import (
	"encoding/json"
	"find/glog"
	"find/source/helpfunc"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"time"
)

type Eth struct {
	Client *rpc.Client
}

func (in *Eth) GetBalance(address string) (balance *big.Int, ok bool, err error) {
RE:
	balance, err = helpfunc.GetBalance(in.Client, address)
	if err != nil {
		er, ok := err.(rpc.HTTPError)
		// 请求过于频繁 需要休眠
		if ok && er.StatusCode == 429 {
			var BodyJson struct {
				Error struct {
					Data struct {
						Rate struct {
							BackoffSeconds int `json:"backoff_seconds"`
						}
					}
				}
			}
			_ = json.Unmarshal(er.Body, &BodyJson)
			BackoffSeconds := BodyJson.Error.Data.Rate.BackoffSeconds
			// glog.Errorf("休眠 %d 秒", BackoffSeconds)
			time.Sleep(time.Duration(BackoffSeconds) * time.Second)
			goto RE
		}
		glog.ErrorExist(`eth: 获取余额出现未知错误: `+ err.Error())
		time.Sleep(time.Second * 60) // 任何错误都睡一分钟继续?
		goto RE
		// return nil, false, errors.New("eth获取余额错误 e.Call : " + err.Error())
	}

	if balance.Cmp(big.NewInt(0)) != 1 { // 1e16 暂时先放宽要求 有余额就录入
		return balance, false, nil
	}

	return balance, true, nil
}
