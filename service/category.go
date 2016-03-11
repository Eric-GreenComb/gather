package service

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/category/client/thrift"
	thriftservice "github.com/banerwai/micros/category/service"
	thriftcategory "github.com/banerwai/micros/category/thrift/gen-go/category"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/micros/common/etcd"
)

type CategoryService struct {
	trans thrift.TTransport
}

func (self *CategoryService) DefaultService() (thriftservice.CategoryService, error) {
	_addr, _err := etcd.GetValue("/banerwai/micros/category/addr")

	if _err != nil {
		return nil, _err
	}

	return self.OpenService(_addr)
}

func (self *CategoryService) OpenService(addr string) (thriftservice.CategoryService, error) {

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
	cli := thriftcategory.NewCategoryServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.CategoryService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *CategoryService) CloseService() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
