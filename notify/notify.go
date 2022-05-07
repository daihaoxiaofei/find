package notify

import (
	"find/config"
	"sync"
)

var notifys []iNotify

type iNotify interface {
	PushMsg(sprintf string)
}

func PushMsg(msg string) {
	for _, in := range getNotifys() {
		in.PushMsg(msg)
	}
}

var once sync.Once

// 懒汉单例模式
func getNotifys() []iNotify {
	once.Do(func() {
		if config.Notify.DD.Open {
			notifys = append(notifys, &ddNotify{})
		}
		if config.Notify.FeiShu.Open {
			notifys = append(notifys, &feiShuNotify{})
		}
		if config.Notify.Email.Open {
			notifys = append(notifys, &emailNotify{})
		}
	})
	return notifys
}
