package glog

import (
	"find/config"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	LogPath = `log`
	once    sync.Once
	gl      = &gLog{}
)

type logMod uint16

const (
	// 日志模式
	ModeDiscard logMod = 1 << iota
	ModeFile
	ModeStdout
	ModeSF = ModeFile | ModeStdout
)

type gLog struct {
	err  *log.Logger
	info *log.Logger
	mu   sync.Mutex
}

type eMap struct {
	Map map[string]int64
	mu  *sync.Mutex
}

var em = &eMap{
	Map: make(map[string]int64),
	mu:  &sync.Mutex{},
}

// var EMap sync.Map

func ErrorExist(s string) {
	em.mu.Lock()
	defer em.mu.Unlock()
	n, ok := em.Map[s]
	if !ok {
		n = 0
		l().err.Output(2, s)
	}
	n = n + 1
	em.Map[s] = n
	Info(s, n)
}

func Error(v ...interface{}) {
	l().err.Output(2, fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	l().err.Output(2, fmt.Sprintf(format, v...))
}
func Info(v ...interface{}) {
	l().info.Output(2, fmt.Sprintln(v...))
}
func Infof(format string, v ...interface{}) {
	l().info.Output(2, fmt.Sprintf(format, v...))
}

// l 这种设计是为了解决配置文件路径在测试环境下也可以保持位置
// 凡是需要加载文件位置之类随环境变化的项目都使用了这种懒汉单例模式
// 允许在功能使用前对其进行设置
// 后来觉得这种模式也不是最优解  参考本项目中config模块设计更加简洁
// 这个就暂时不改了 留一种多样性
func l() *gLog {
	once.Do(func() {
		// prefix:=`[31m[error][0m `
		gl.err = log.New(ioutil.Discard, ``, log.LstdFlags|log.Llongfile)
		gl.info = log.New(ioutil.Discard, ``, log.LstdFlags|log.Lshortfile)

		set(`err`, logMod(config.Run.ErrLogMod))
		set(`info`, logMod(config.Run.InfoLogMod))
	})

	return gl
}

func set(name string, mod logMod) {
	gl.mu.Lock()
	defer gl.mu.Unlock()

	var l *log.Logger
	switch name {
	case `err`:
		l = gl.err
	case `info`:
		l = gl.info
	default:
		panic(`无效的name` + name)
	}

	if mod&ModeDiscard != 0 {
		l.SetOutput(ioutil.Discard)
	} else {
		var allWriters []io.Writer
		if mod&ModeFile != 0 {
			if _, err := os.Stat(LogPath); err != nil {
				err := os.Mkdir(LogPath, os.ModePerm)
				if err != nil {
					panic(`路径创建失败 Mkdir err: ` + err.Error())
				}
			}
			File, err := os.OpenFile(filepath.Join(LogPath, name+`.log`), os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(`打开日志文件时出现错误:` + err.Error())
			}
			allWriters = append(allWriters, File)
		}
		if mod&ModeStdout != 0 {
			allWriters = append(allWriters, os.Stdout)
		}
		l.SetOutput(io.MultiWriter(allWriters...))
	}
}
