package notify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"find/pkg/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type ddNotify struct {
}

func (f *ddNotify) PushMsg(msg string) {
	// doc: https://open.dingtalk.com/document/robots/customize-robot-security-settings
	data := map[string]interface{}{
		`msgtype`: `text`,
		`text`:    map[string]string{"content": msg},
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Printf("Do-构造参数失败 %+v", err)
		return
	}
	Sign, timestamp := toSign(config.Notify.DD.Secret)
	// 发起请求
	resp, err := http.Post(config.Notify.DD.Url+`&timestamp=`+timestamp+`&sign=`+Sign, "application/json;UTF-8", bytes.NewReader(dataJson))
	if err != nil {
		log.Printf("请求失败, %+v", err)
		return
	}
	// 关闭http流
	defer resp.Body.Close()

	// 读取返回
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取失败 %+v", err)
		return
	}
	fmt.Println(string(body))
	return
}

func toSign(secret string) (sign string, timestamp string) {
	if secret == "" {
		return "", ""
	}
	timestamp = strconv.FormatInt(time.Now().Unix()*1e3, 10)
	signData := computeHmacSha256(timestamp+"\n"+secret, secret)
	return url.QueryEscape(signData), timestamp
}

func computeHmacSha256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
