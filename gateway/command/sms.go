package command

import (
	"encoding/json"
	gatherbean "github.com/banerwai/gather/bean"

	gatherredis "github.com/banerwai/gather/common/redis"

	"github.com/banerwai/gather/common/gearman"
	gearmanclient "github.com/mikespook/gearman-go/client"
)

type SmsService struct {
}

func (self *SmsService) LPush2Redis(key string, sms gatherbean.SMS) error {
	b, err := json.Marshal(sms)
	if err != nil {
		return err
	}

	conn := gatherredis.RedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("LPUSH", key, string(b))

	if err != nil {
		return err
	}

	return nil
}

func (self *SmsService) Send2Gearman(key string) error {
	c, err := gearmanclient.New(gearmanclient.Network, gearman.GearmanAddr)
	if err != nil {
		return err
	}
	defer c.Close()

	_, err = c.DoBg("SendSms", []byte(key), gearmanclient.JobNormal)
	if err != nil {
		return err
	}

	return nil
}
