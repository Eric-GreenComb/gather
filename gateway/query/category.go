package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/category/client/thrift"
	thriftservice "github.com/banerwai/micros/query/category/service"
	thriftcategory "github.com/banerwai/micros/query/category/thrift/gen-go/category"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type CategoryService struct {
	trans thrift.TTransport
	addr  string
}

func (self *CategoryService) Default() (thriftservice.CategoryService, error) {

	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *CategoryService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosQueryCategory)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("category query micro service is 0")
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *CategoryService) Open() (thriftservice.CategoryService, error) {

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
	cli := thriftcategory.NewCategoryServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.CategoryService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *CategoryService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
