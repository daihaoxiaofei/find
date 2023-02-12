package notify

import (
	"find/pkg/config"
	"sync"
)

var notifyArr []iNotify

type iNotify interface {
	PushMsg(sprintf string)
}

func PushMsg(msg string) {
	for _, in := range getNotifyArr() {
		in.PushMsg(msg)
	}
}

var once sync.Once

// 懒汉单例模式  用到时再加载
func getNotifyArr() []iNotify {
	once.Do(func() {
		if config.Notify.DD.Open {
			notifyArr = append(notifyArr, &ddNotify{})
		}
		if config.Notify.FeiShu.Open {
			notifyArr = append(notifyArr, &feiShuNotify{})
		}
		if config.Notify.Email.Open {
			notifyArr = append(notifyArr, &emailNotify{})
		}
	})
	return notifyArr
}
