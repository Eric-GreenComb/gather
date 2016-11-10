package command

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/command/workhistory/client/thrift"
	thriftservice "github.com/banerwai/micros/command/workhistory/service"
	thriftworkhistory "github.com/banerwai/micros/command/workhistory/thrift/gen-go/workhistory"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// WorkHistoryService WorkHistoryService
type WorkHistoryService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (whs *WorkHistoryService) Default() (thriftservice.WorkHistoryService, error) {
	_err := whs.Init()
	if _err != nil {
		return nil, _err
	}

	return whs.Open()
}

// Init service get addr
func (whs *WorkHistoryService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosCommandWorkhistory)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("workhistory command micro service is 0")
	}

	whs.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

// Open service open addr
func (whs *WorkHistoryService) Open() (thriftservice.WorkHistoryService, error) {

	transportSocket, err := thrift.NewTSocket(whs.addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	whs.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := whs.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thriftworkhistory.NewWorkHistoryServiceClientFactory(whs.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.WorkHistoryService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (whs *WorkHistoryService) Close() {
	if whs.trans.IsOpen() {
		whs.trans.Close()
	}
}
