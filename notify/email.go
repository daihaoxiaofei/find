package notify

import (
	"find/config"
	"find/glog"
	"gopkg.in/gomail.v2"
)

type emailNotify struct{}

func (f *emailNotify) PushMsg(msg string) {
	Cf := config.Notify.Email
	Name := `find程序通知`
	m := gomail.NewMessage()
	m.SetHeader("From", Cf.From)                        // 发件人
	m.SetHeader("From", m.FormatAddress(Cf.From, Name)) // 别名
	m.SetHeader("To", Cf.To)                            // 收件人
	m.SetHeader("Subject", "["+Name+"]通知邮件")
	m.SetBody("text/plain", msg)

	d := gomail.NewDialer(Cf.HOST, Cf.PORT, Cf.From, Cf.Pwd)
	if err := d.DialAndSend(m); err != nil {
		glog.Error(`发送邮件时出现错误err: ` + err.Error())
	}
}
