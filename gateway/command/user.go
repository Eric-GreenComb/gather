package command

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/user/client/thrift"
	thriftservice "github.com/banerwai/micros/command/user/service"
	thriftuser "github.com/banerwai/micros/command/user/thrift/gen-go/user"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type UserService struct {
	trans thrift.TTransport
	addr  string
}

func (self *UserService) Default() (thriftservice.UserService, error) {
	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *UserService) Init() error {

	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_COMMAND_USER)

	if _err != nil {
		return _err
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *UserService) Open() (thriftservice.UserService, error) {

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
	cli := thriftuser.NewUserServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.UserService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *UserService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
