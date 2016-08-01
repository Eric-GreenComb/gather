package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/account/client/thrift"
	thriftservice "github.com/banerwai/micros/query/account/service"
	thriftaccount "github.com/banerwai/micros/query/account/thrift/gen-go/account"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type AccountService struct {
	trans thrift.TTransport
	addr  string
}

func (self *AccountService) Default() (thriftservice.AccountService, error) {

	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *AccountService) Init() error {

	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_QUERY_ACCOUNT)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("profile query micro service is 0")
	}
	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]
	return nil
}

func (self *AccountService) Open() (thriftservice.AccountService, error) {

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
	cli := thriftaccount.NewAccountServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.AccountService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *AccountService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
