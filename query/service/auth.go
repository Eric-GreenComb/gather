package service

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/auth/client/thrift"
	thriftservice "github.com/banerwai/micros/query/auth/service"
	thriftauth "github.com/banerwai/micros/query/auth/thrift/gen-go/auth"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type AuthService struct {
	trans thrift.TTransport
	addr  string
}

func (self *AuthService) Default() (thriftservice.AuthService, error) {
	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *AuthService) Init() error {

	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_QUERY_AUTH)

	if _err != nil {
		return _err
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *AuthService) Open() (thriftservice.AuthService, error) {

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
	cli := thriftauth.NewAuthServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.AuthService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *AuthService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
