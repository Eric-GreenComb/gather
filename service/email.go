package service

import (
	"encoding/json"
	"flag"
	"github.com/banerwai/global/bean"
	"github.com/nats-io/nats"
)

var (
	NatsUrls string
)

func init() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	flag.Parse()
	NatsUrls = *urls
}

type EmailService struct {
}

func (self *EmailService) SendEmail(email bean.Email) error {
	nc, err := nats.Connect(NatsUrls)
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

func (self *EmailService) SendTpl(email_extra bean.EmailExtra) error {
	nc, err := nats.Connect(NatsUrls)
	defer nc.Close()

	if err != nil {
		return err
	}

	b, err := json.Marshal(email_extra)
	if err != nil {
		return err
	}

	nc.Publish("tpl", b)
	nc.Flush()

	return nil
}
