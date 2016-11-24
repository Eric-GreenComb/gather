package query

import (
	"errors"
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/contact/client/thrift"
	thriftservice "github.com/banerwai/micros/query/contact/service"
	thriftcontact "github.com/banerwai/micros/query/contact/thrift/gen-go/contact"

	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// ContactService ContactServices
type ContactService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (cs *ContactService) Default() (thriftservice.ContactService, error) {

	_err := cs.Init()
	if _err != nil {
		return nil, _err
	}

	return cs.Open()
}

// Init service get addr
func (cs *ContactService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosQueryContact)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("profile query micro service is 0")
	}
	cs.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]
	return nil
}

// Open service open addr
func (cs *ContactService) Open() (thriftservice.ContactService, error) {

	transportSocket, err := thrift.NewTSocket(cs.addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	cs.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := cs.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thriftcontact.NewContactServiceClientFactory(cs.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.ContactService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (cs *ContactService) Close() {
	if cs.trans.IsOpen() {
		cs.trans.Close()
	}
}
