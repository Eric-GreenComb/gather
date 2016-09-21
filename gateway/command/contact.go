package command

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/contact/client/thrift"
	thriftservice "github.com/banerwai/micros/command/contact/service"
	thriftcontact "github.com/banerwai/micros/command/contact/thrift/gen-go/contact"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type ContactService struct {
	trans thrift.TTransport
	addr  string
}

func (self *ContactService) Default() (thriftservice.ContactService, error) {
	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *ContactService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosCommandContact)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("contact command micro service is 0")
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *ContactService) Open() (thriftservice.ContactService, error) {

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
	cli := thriftcontact.NewContactServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.ContactService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *ContactService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
