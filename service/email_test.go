package service

import (
	"testing"

	thriftemailbean "github.com/banerwai/micros/email/thrift/gen-go/email"
)

// need start micro render service localhost:6003
func TestEmailDefaultService(t *testing.T) {

	var _email EmailService
	__email_service, _ := _email.DefaultService()
	defer _email.CloseService()

	_email_bean := thriftemailbean.Email{}
	// change email
	_email_bean.Host = "smtp.126.com:25"
	_email_bean.User = "xxx@126.com"
	_email_bean.Password = "xxx"
	_email_bean.To = "xxx@126.com"
	_email_bean.Subject = "Hi"
	_email_bean.Body = "This is a test email"
	_email_bean.Mailtype = "html"

	v := __email_service.SendEmail(&_email_bean)

	if v != true {
		t.Errorf("TestCategoryDefaultService error")
	}
}
