package command

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/profile/client/thrift"
	thriftservice "github.com/banerwai/micros/command/profile/service"
	thriftprofile "github.com/banerwai/micros/command/profile/thrift/gen-go/profile"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type ProfileService struct {
	trans thrift.TTransport
	addr  string
}

func (self *ProfileService) Default() (thriftservice.ProfileService, error) {
	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *ProfileService) Init() error {

	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_COMMAND_PROFILE)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("profile command micro service is 0")
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *ProfileService) Open() (thriftservice.ProfileService, error) {

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
	cli := thriftprofile.NewProfileServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.ProfileService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *ProfileService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
