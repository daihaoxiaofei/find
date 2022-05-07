package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Run    *run
	Source *source
	Notify *notify
)

type run struct {
	NumLine    int    // 每个客户端线程数
	ErrLogMod  uint16 // 错误日志模式   模式 1:安静 2:写入文件 4:标准输出 6:写入文件和标准输出
	InfoLogMod uint16 // 日常日志模式
}

type source struct {
	Urls struct {
		Eth  []string
		Bsc  []string
		HECO []string
		OKEX []string
	}
}

type notify struct {
	DD struct {
		Open   bool
		Url    string
		Secret string
	}
	FeiShu struct {
		Open bool
		Url  string
	}
	Email struct {
		Open bool
		HOST string
		PORT int
		From string
		Pwd  string
		To   string
	}
}

func init() {
	vp := viper.New()

	if (runtime.GOOS == `linux` && os.Args[0][len(os.Args[0])-5:] == `.test`) ||
		runtime.GOOS == `windows` && filepath.Base(os.Args[0])[:7] == `___Test` {
		fmt.Println(runtime.GOOS, `测试环境配置文件`)
		_, onPath, _, _ := runtime.Caller(0)
		onDir := filepath.Dir(onPath)
		vp.AddConfigPath(onDir)
	} else {
		vp.AddConfigPath("config")
	}


	vp.SetConfigType("yaml")
	vp.SetConfigName("config")
	err := vp.ReadInConfig()
	if err != nil {
		panic(`读取配置文件时出现错误:` + err.Error())
	}

	err = vp.UnmarshalKey("Run", &Run)
	if err != nil {
		panic(`读取配置文件时出现错误:` + err.Error())
	}
	err = vp.UnmarshalKey("Source", &Source)
	if err != nil {
		panic(`读取配置文件时出现错误:` + err.Error())
	}
	err = vp.UnmarshalKey("Notify", &Notify)
	if err != nil {
		panic(`读取配置文件时出现错误:` + err.Error())
	}
}
