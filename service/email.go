package service

import (
	"github.com/banerwai/gather/gateway/command"

	"github.com/banerwai/global/bean"
)

type EmailService struct {
}

func (self *EmailService) SendEmail(key string, email bean.Email) error {
	var _service command.EmailService
	_err := _service.LPush2Redis(key, email)
	if _err != nil {
		return _err
	}
	_err = _service.Send2Gearman(key)
	if _err != nil {
		return _err
	}
	return nil
}
