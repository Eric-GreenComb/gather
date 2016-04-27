package service

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/render/client/thrift"
	thriftservice "github.com/banerwai/micros/query/render/service"
	thriftrender "github.com/banerwai/micros/query/render/thrift/gen-go/render"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type RenderService struct {
	trans thrift.TTransport
	addr  string
}

func (self *RenderService) Default() (thriftservice.RenderService, error) {
	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *RenderService) Init() error {
	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_QUERY_RENDER)

	if _err != nil {
		return _err
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *RenderService) Open() (thriftservice.RenderService, error) {

	transportSocket, err := thrift.NewTSocket(self.addr)
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

func (self *RenderService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
