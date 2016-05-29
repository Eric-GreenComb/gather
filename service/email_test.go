package service

import (
	"testing"

	"github.com/banerwai/global/bean"
)

// func TestEmailServiceSendEmail(t *testing.T) {

// 	var _email EmailService

// 	_email_bean := bean.Email{}
// 	// change email
// 	_email_bean.Host = "smtp.126.com:25"
// 	_email_bean.User = "xxx@126.com"
// 	_email_bean.Password = "xxx"
// 	_email_bean.To = "xxx@126.com"
// 	_email_bean.Subject = "Hi"
// 	_email_bean.Body = "This is a test email"
// 	_email_bean.Mailtype = "html"

// 	_err := _email.SendEmail(_email_bean)

// 	if _err != nil {
// 		t.Errorf("SendEmail error")
// 	}
// }

func TestEmailServiceSendTpl(t *testing.T) {
	var _service EmailService

	_email_extra := bean.EmailExtra{}
	_email_extra.Email.Host = "smtp.126.com:25"
	_email_extra.Email.User = "ministor@126.com"
	_email_extra.Email.Password = "xxxxxx"
	_email_extra.Email.To = "ministor@126.com"
	_email_extra.Email.Subject = "this is a tpl email"
	_email_extra.Email.Mailtype = "html"

	_email_extra.TempName = "hello"

	_map_parse := make(map[string]string)
	_map_parse["Hi"] = "Hello"
	_map_parse["Name"] = "Eric"
	_email_extra.Parse = _map_parse

	_err := _service.SendTpl(_email_extra)

	if _err != nil {
		t.Errorf("SendTpl error")
	}
}
