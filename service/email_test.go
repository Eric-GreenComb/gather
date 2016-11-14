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

	_emailExtra := bean.EmailExtra{}
	_emailExtra.Email.Host = "smtp.126.com:25"
	_emailExtra.Email.User = "ministor@126.com"
	_emailExtra.Email.Password = "xxxxxx"
	_emailExtra.Email.To = "ministor@126.com"
	_emailExtra.Email.Subject = "this is a tpl email"
	_emailExtra.Email.Mailtype = "html"

	_emailExtra.TempName = "hello"

	_mapParse := make(map[string]string)
	_mapParse["Hi"] = "Hello"
	_mapParse["Name"] = "Eric"
	_emailExtra.Parse = _mapParse

	_err := _service.SendTpl(_emailExtra)

	if _err != nil {
		t.Errorf("SendTpl error")
	}
}
