package service

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/render/client/thrift"
	thriftservice "github.com/banerwai/micros/query/render/service"
	thriftrender "github.com/banerwai/micros/query/render/thrift/gen-go/render"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/micros/common/etcd"
)

type RenderService struct {
	trans thrift.TTransport
}

func (self *RenderService) DefaultService() (thriftservice.RenderService, error) {
	_addr, _err := etcd.GetValue("/banerwai/micros/query/render/addr")

	if _err != nil {
		return nil, _err
	}

	return self.OpenService(_addr)
}

func (self *RenderService) OpenService(addr string) (thriftservice.RenderService, error) {

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
	cli := thriftrender.NewRenderServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.RenderService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *RenderService) CloseService() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
