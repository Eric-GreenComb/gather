package command

import (
	"testing"

	"github.com/banerwai/global/bean"
)

func TestEmailLPush2Redis(t *testing.T) {

	var _email EmailService

	_email_bean := bean.Email{}
	// change email
	_email_bean.Host = "smtp.126.com:25"
	_email_bean.User = "xxx@126.com"
	_email_bean.Password = "xxx"
	_email_bean.To = "xxx@126.com"
	_email_bean.Subject = "Hi"
	_email_bean.Body = "This is a test email"
	_email_bean.Mailtype = "html"

	_err := _email.LPush2Redis("banerwai:email:activeuser", _email_bean)

	if _err != nil {
		t.Errorf("LPush2Redis error")
	}
}

func TestEmailSend2Gearman(t *testing.T) {

	var _email EmailService

	_err := _email.Send2Gearman("banerwai:email:activeuser")

	if _err != nil {
		t.Errorf("Send2Gearman error")
	}
}
