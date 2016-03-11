package service

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/email/client/thrift"
	thriftservice "github.com/banerwai/micros/email/service"
	thriftemail "github.com/banerwai/micros/email/thrift/gen-go/email"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/micros/common/etcd"
)

type EmailService struct {
	trans thrift.TTransport
}

func (self *EmailService) DefaultService() (thriftservice.EmailService, error) {
	_addr, _err := etcd.GetValue("/banerwai/micros/email/addr")

	if _err != nil {
		return nil, _err
	}

	return self.OpenService(_addr)
}

func (self *EmailService) OpenService(addr string) (thriftservice.EmailService, error) {

	transportSocket, err := thrift.NewTSocket(addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	self.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := self.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thriftemail.NewEmailServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.EmailService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *EmailService) CloseService() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
