package service

import (
	"os"

	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/token/client/thrift"
	thriftservice "github.com/banerwai/micros/token/service"
	thrifttoken "github.com/banerwai/micros/token/thrift/gen-go/token"

	gatherthrift "github.com/banerwai/gather/common/thrift"
)

type TokenService struct {
	trans thrift.TTransport
}

func (self *TokenService) OpenService(addr string) thriftservice.TokenService {

	transportSocket, err := thrift.NewTSocket(addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		os.Exit(1)
	}
	self.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := self.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		os.Exit(1)
	}
	cli := thrifttoken.NewTokenServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.TokenService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc
}

func (self *TokenService) CloseService() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
