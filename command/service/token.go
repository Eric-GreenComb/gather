package service

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/token/client/thrift"
	thriftservice "github.com/banerwai/micros/command/token/service"
	thrifttoken "github.com/banerwai/micros/command/token/thrift/gen-go/token"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/micros/common/etcd"
)

type TokenService struct {
	trans thrift.TTransport
}

func (self *TokenService) DefaultService() (thriftservice.TokenService, error) {
	_addr, _err := etcd.GetValue("/banerwai/micros/command/token/addr")

	if _err != nil {
		return nil, _err
	}

	return self.OpenService(_addr)
}

func (self *TokenService) OpenService(addr string) (thriftservice.TokenService, error) {

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
	cli := thrifttoken.NewTokenServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.TokenService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *TokenService) CloseService() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
