package command

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/user/client/thrift"
	thriftservice "github.com/banerwai/micros/command/user/service"
	thriftuser "github.com/banerwai/micros/command/user/thrift/gen-go/user"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// UserService UserService
type UserService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (us *UserService) Default() (thriftservice.UserService, error) {
	_err := us.Init()
	if _err != nil {
		return nil, _err
	}

	return us.Open()
}

// Init service get addr
func (us *UserService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosCommandUser)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("user command micro service is 0")
	}

	us.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

// Open service open addr
func (us *UserService) Open() (thriftservice.UserService, error) {

	transportSocket, err := thrift.NewTSocket(us.addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	us.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := us.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thriftuser.NewUserServiceClientFactory(us.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.UserService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (us *UserService) Close() {
	if us.trans.IsOpen() {
		us.trans.Close()
	}
}
