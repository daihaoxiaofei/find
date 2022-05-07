package keypair

import (
	"crypto/ecdsa"
	"find/glog"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Pair struct {
	Private string
	Address string
}

// CreateKey 生成私钥和地址
// 以太坊的私钥、公钥、地址之间的关系是，随机生成私钥、根据私钥计算出公钥、根据公钥计算出地址。
// 私钥是256bit，相当于32字节；用16进制表示的话，就是64个字符。
// 私钥的生成，本质上和选取一个1到2^256之间的数字几乎一致，其中一种方法就是在很大的空间中选择一个随机数，然后使用SHA256计算其哈希，作为私钥。
// 通过私钥，可以生成一个64字节的公钥，生成办法是通过椭圆曲线算法，这个算法是确定性的。
// 私钥进行Keccak-256计算之后，用16进制表示，保留最后20位，就是地址了。
func CreateKey() <-chan Pair {
	Ad := make(chan Pair, 10)
	go func() {
		n := 0
		for {
			// 创建私钥
			privateKey, err := crypto.GenerateKey()
			if err != nil {
				glog.Errorf("创建私钥失败 crypto.GenerateKey()  %v", err)
			}
			/*	//可通过此代码导入私钥
				privateKey,err=crypto.HexToECDSA("93d5d04256882aaad507ff09f510969f347758109793448aa79e1b4dbe5f6efa")
				if err != nil {
					log.Fatal(err)
				}
			*/
			privateKeyBytes := crypto.FromECDSA(privateKey)
			private := hexutil.Encode(privateKeyBytes)[2:]
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				glog.Error("publicKeyECDSA 生成失败")
			}
			address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
			Ad <- Pair{
				Private: private,
				Address: address,
			}
			n++
			if n%1000 == 0 {
				glog.Infof("已发布 %d 条地址", n)
			}
		}
	}()

	return Ad
}
