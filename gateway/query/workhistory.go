package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/workhistory/client/thrift"
	thriftservice "github.com/banerwai/micros/query/workhistory/service"
	thriftworkhistory "github.com/banerwai/micros/query/workhistory/thrift/gen-go/workhistory"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type WorkHistoryService struct {
	trans thrift.TTransport
	addr  string
}

func (self *WorkHistoryService) Default() (thriftservice.WorkHistoryService, error) {

	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *WorkHistoryService) Init() error {

	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_QUERY_WORKHISTORY)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("workhistory query micro service is 0")
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *WorkHistoryService) Open() (thriftservice.WorkHistoryService, error) {

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
	cli := thriftworkhistory.NewWorkHistoryServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.WorkHistoryService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *WorkHistoryService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
