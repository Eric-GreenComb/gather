package service

import (
	"encoding/json"
	"github.com/banerwai/gather/common/flagparse"
	"github.com/banerwai/global/bean"
	"github.com/nats-io/nats"
)

// EmailService EmailService
type EmailService struct {
}

// SendEmail send email bean
func (es *EmailService) SendEmail(email bean.Email) error {
	nc, err := nats.Connect(flagparse.NatsUrls)
	defer nc.Close()

	if err != nil {
		return err
	}

	b, err := json.Marshal(email)
	if err != nil {
		return err
	}

	nc.Publish("mail", b)
	nc.Flush()

	return nil
}

// SendTpl send tpl
func (es *EmailService) SendTpl(emailEtra bean.EmailExtra) error {
	nc, err := nats.Connect(flagparse.NatsUrls)
	defer nc.Close()

	if err != nil {
		return err
	}

	b, err := json.Marshal(emailEtra)
	if err != nil {
		return err
	}

	nc.Publish("tpl", b)
	nc.Flush()

	return nil
}
