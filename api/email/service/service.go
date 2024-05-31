package service

import (
	"gopkg.in/gomail.v2"
	"gwf/api/email/mod"
)

type Mailer struct {
	mod.EmailBody
}

func (ctx *Mailer) SendEmail() error {
	// TODO
	msg := gomail.NewMessage()
	msg.SetHeader("From", ctx.F)
	msg.SetHeader("To", ctx.T)
	msg.SetHeader("Cc", ctx.C)
	msg.SetHeader("Subject", "账号申请")
	msg.SetBody("text/html", ctx.HtmlBody)
	dialer := gomail.NewDialer("smtp.qq.com", 465, ctx.Account, ctx.Password)
	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}
