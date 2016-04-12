package service

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/profile/client/thrift"
	thriftservice "github.com/banerwai/micros/query/profile/service"
	thriftprofile "github.com/banerwai/micros/query/profile/thrift/gen-go/profile"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/micros/common/etcd"
)

type ProfileService struct {
	trans thrift.TTransport
}

func (self *CategoryService) DefaultService() (thriftservice, error) {
	_addr, _err := etcd.GetValue("/banerwai/micros/query/profile/addr")

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
