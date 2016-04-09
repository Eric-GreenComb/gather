package service

import (
	"testing"

	"github.com/banerwai/gather/bean"
)

func TestSmsLPush2Redis(t *testing.T) {

	var _sms_service SmsService

	_sms_bean := bean.Sms{}
	_sms_bean.Name = "xxxxxxxx"
	_sms_bean.Pwd = "xxxxxxx"
	_sms_bean.Content = "register info : xxx"
	_sms_bean.Mobile = "13811111111,13911111111"
	_sms_bean.Sign = "banerwaiSign"
	_sms_bean.Extno = "none"

	_err := _sms_service.LPush2Redis("banerwai:sms:activeuser", _sms_bean)

	if _err != nil {
		t.Errorf("SmsLPush2Redis error")
	}
}
