package service

import (
	"github.com/banerwai/gather/gateway/command"

	"github.com/banerwai/global/bean"
)

type SmsService struct {
}

func (self *SmsService) SendSms(key string, sms bean.SMS) error {
	var _service command.SmsService
	_err := _service.LPush2Redis(key, sms)
	if _err != nil {
		return _err
	}
	_err = _service.Send2Gearman(key)
	if _err != nil {
		return _err
	}
	return nil
}
