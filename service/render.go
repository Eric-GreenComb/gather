package service

import (
	"os"

	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/render/client/thrift"
	thriftservice "github.com/banerwai/micros/render/service"
	thriftrender "github.com/banerwai/micros/render/thrift/gen-go/render"

	gatherthrift "github.com/banerwai/gather/common/thrift"
)

type RenderService struct {
	trans thrift.TTransport
}

func (self *RenderService) OpenService(addr string) thriftservice.RenderService {

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
	cli := thriftrender.NewRenderServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.RenderService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc
}

func (self *RenderService) CloseService() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
