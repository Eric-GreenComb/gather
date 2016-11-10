package command

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/profile/client/thrift"
	thriftservice "github.com/banerwai/micros/command/profile/service"
	thriftprofile "github.com/banerwai/micros/command/profile/thrift/gen-go/profile"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// ProfileService ProfileService
type ProfileService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (ps *ProfileService) Default() (thriftservice.ProfileService, error) {
	_err := ps.Init()
	if _err != nil {
		return nil, _err
	}

	return ps.Open()
}

// Init service get addr
func (ps *ProfileService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosCommandProfile)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("profile command micro service is 0")
	}

	ps.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

// Open service open addr
func (ps *ProfileService) Open() (thriftservice.ProfileService, error) {

	transportSocket, err := thrift.NewTSocket(ps.addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	ps.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := ps.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thriftprofile.NewProfileServiceClientFactory(ps.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.ProfileService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (ps *ProfileService) Close() {
	if ps.trans.IsOpen() {
		ps.trans.Close()
	}
}
