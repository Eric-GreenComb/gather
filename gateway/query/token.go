package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/token/client/thrift"
	thriftservice "github.com/banerwai/micros/query/token/service"
	thrifttoken "github.com/banerwai/micros/query/token/thrift/gen-go/token"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type TokenService struct {
	trans thrift.TTransport
	addr  string
}

func (self *TokenService) Default() (thriftservice.TokenService, error) {
	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *TokenService) Init() error {

	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_QUERY_TOKEN)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("token query micro service is 0")
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *TokenService) Open() (thriftservice.TokenService, error) {

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
	cli := thrifttoken.NewTokenServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.TokenService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *TokenService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
