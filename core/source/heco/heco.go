package heco

import (
	"find/core/keypair"
	"find/core/source/helpfunc"
	"find/pkg/glog"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"go.uber.org/zap"
	"net/url"
)

type Client struct {
	cli      *rpc.Client
	from     string
	initPNum int
}

func NewClient(url string, initPNum int) *Client {
	c, err := rpc.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("连接%s,错误 rpc.Dial %#v", url, err))
	}
	return &Client{
		cli:      c,
		from:     `heco ` + url,
		initPNum: initPNum,
	}
}

func (c *Client) Find(ad <-chan keypair.Pair) {
	errFun := func(err error) {
		rpcHTTPError, ok := err.(rpc.HTTPError)
		if ok {
			glog.Debug(`errFun`, zap.String(`rpcHTTPError`, fmt.Sprintf("%#v\n", rpcHTTPError)))
			return
		}

		urlError, ok := err.(*url.Error)
		if ok {
			glog.Debug(`errFun`, zap.String(`urlError`, fmt.Sprintf("%#v\n", urlError)))
			return
		}
		glog.Error(c.from+" Cli.Call 未知错误:", zap.String(`err`, fmt.Sprintf("%#v\n", err)))
	}
	helpfunc.Find(c.cli, ad, c.from, c.initPNum, errFun)
}
