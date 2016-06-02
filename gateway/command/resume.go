package command

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/resume/client/thrift"
	thriftservice "github.com/banerwai/micros/command/resume/service"
	thriftresume "github.com/banerwai/micros/command/resume/thrift/gen-go/resume"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	banerwaiglobal "github.com/banerwai/global"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

type ResumeService struct {
	trans thrift.TTransport
	addr  string
}

func (self *ResumeService) Default() (thriftservice.ResumeService, error) {
	_err := self.Init()
	if _err != nil {
		return nil, _err
	}

	return self.Open()
}

func (self *ResumeService) Init() error {

	_addrs, _err := etcd.GetServicesByName(banerwaiglobal.ETCD_KEY_MICROS_COMMAND_RESUME)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("resume command micro service is 0")
	}

	self.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

func (self *ResumeService) Open() (thriftservice.ResumeService, error) {

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
	cli := thriftresume.NewResumeServiceClientFactory(self.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.ResumeService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

func (self *ResumeService) Close() {
	if self.trans.IsOpen() {
		self.trans.Close()
	}
}
