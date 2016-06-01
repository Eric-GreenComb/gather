package service

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/nats-io/nats"
)

type SmsService struct {
}

func (self *SmsService) SendSms(sms bean.SMS) error {
	nc, err := nats.Connect(NatsUrls)
	defer nc.Close()

	if err != nil {
		return err
	}

	b, err := json.Marshal(sms)
	if err != nil {
		return err
	}

	nc.Publish("sms", b)
	nc.Flush()

	return nil
}
