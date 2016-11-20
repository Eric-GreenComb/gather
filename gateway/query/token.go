package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/token/client/thrift"
	thriftservice "github.com/banerwai/micros/query/token/service"
	thrifttoken "github.com/banerwai/micros/query/token/thrift/gen-go/token"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// TokenService TokenService
type TokenService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (ts *TokenService) Default() (thriftservice.TokenService, error) {
	_err := ts.Init()
	if _err != nil {
		return nil, _err
	}

	return ts.Open()
}

// Init service get addr
func (ts *TokenService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosQueryToken)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("token query micro service is 0")
	}

	ts.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

// Open service open addr
func (ts *TokenService) Open() (thriftservice.TokenService, error) {

	transportSocket, err := thrift.NewTSocket(ts.addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	ts.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := ts.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thrifttoken.NewTokenServiceClientFactory(ts.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.TokenService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (ts *TokenService) Close() {
	if ts.trans.IsOpen() {
		ts.trans.Close()
	}
}
