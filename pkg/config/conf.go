package config

import (
	"find/pkg/runvalue"
	"io/ioutil"
	"log"
	"path"

	"gopkg.in/yaml.v3"
)

type run struct {
	LogLevel string `yaml:"LogLevel"` // 日志等级
}
type sourceMsg struct {
	InitPNum int      `yaml:"InitPNum"` // 日志等级
	URL      []string `yaml:"URL"`      // 日志等级
}

type source struct {
	ThirdParty sourceMsg `yaml:"ThirdParty"`
	Bsc        sourceMsg `yaml:"Bsc"`
	HECO       sourceMsg `yaml:"HECO"`
	OKEX       sourceMsg `yaml:"OKEX"`
}

type notify struct {
	DD struct {
		Open   bool   `yaml:"Open"`
		Url    string `yaml:"Url"`
		Secret string `yaml:"Secret"`
	} `yaml:"DD"`
	FeiShu struct {
		Open bool   `yaml:"Open"`
		Url  string `yaml:"Url"`
	} `yaml:"FeiShu"`
	Email struct {
		Open bool   `yaml:"Open"`
		HOST string `yaml:"HOST"`
		PORT int    `yaml:"PORT"`
		From string `yaml:"From"`
		Pwd  string `yaml:"Pwd"`
		To   string `yaml:"To"`
	} `yaml:"Email"`
}

var (
	Run    run
	Source source
	Notify notify
)

func init() {
	confPath := path.Join(runvalue.RootPath, "config.yaml")

	// 读取文件所有内容装到 []byte 中
	bytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Panicln(`配置文件读取错误 confPath:`, confPath, err)
	}

	// 调用 Unmarshall 去解码文件内容
	var c struct {
		Source source `yaml:"Source"`
		Notify notify `yaml:"Notify"`
		Run    run    `yaml:"Run"`
	}
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Panicln(`配置文件解析错误`, err)
	}

	Run = c.Run
	Source = c.Source
	Notify = c.Notify
}
