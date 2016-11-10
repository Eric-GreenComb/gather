package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/account/client/thrift"
	thriftservice "github.com/banerwai/micros/query/account/service"
	thriftaccount "github.com/banerwai/micros/query/account/thrift/gen-go/account"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// AccountService AccountService
type AccountService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (as *AccountService) Default() (thriftservice.AccountService, error) {

	_err := as.Init()
	if _err != nil {
		return nil, _err
	}

	return as.Open()
}

// Init service get addr
func (as *AccountService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosQueryAccount)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("profile query micro service is 0")
	}
	as.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]
	return nil
}

// Open service open addr
func (as *AccountService) Open() (thriftservice.AccountService, error) {

	transportSocket, err := thrift.NewTSocket(as.addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	as.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := as.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thriftaccount.NewAccountServiceClientFactory(as.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.AccountService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (as *AccountService) Close() {
	if as.trans.IsOpen() {
		as.trans.Close()
	}
}
