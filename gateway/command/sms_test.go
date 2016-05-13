package command

import (
	"testing"

	"github.com/banerwai/global/bean"
)

func TestSmsLPush2Redis(t *testing.T) {

	var _sms_service SmsService

	_sms_bean := bean.SMS{}
	_sms_bean.RecNum = "13000000000,13000000001"
	_sms_bean.SmsFreeSignName = "NeedChange"
	_sms_bean.SmsTemplateCode = "SMS_585014"
	_sms_bean.SmsParam = "{\"code\":\"1234\",\"product\":\"alidayu\"}"

	_err := _sms_service.LPush2Redis("banerwai:sms:activeuser", _sms_bean)

	if _err != nil {
		t.Errorf("SmsLPush2Redis error")
	}
}

func TestEmailSend2Gearman(t *testing.T) {

	var _sms_service SmsService

	_err := _sms_service.Send2Gearman("banerwai:sms:activeuser")

	if _err != nil {
		t.Errorf("Send2Gearman error")
	}
}
