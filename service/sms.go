package service

import (
	"encoding/json"
	"github.com/banerwai/gather/common/flagparse"
	"github.com/banerwai/global/bean"
	"github.com/nats-io/nats"
)

// SmsService SmsService
type SmsService struct {
}

// SendSms send sms bean
func (ss *SmsService) SendSms(sms bean.SMS) error {
	nc, err := nats.Connect(flagparse.NatsUrls)
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
