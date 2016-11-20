package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/resume/client/thrift"
	thriftservice "github.com/banerwai/micros/query/resume/service"
	thriftresume "github.com/banerwai/micros/query/resume/thrift/gen-go/resume"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// ResumeService ResumeService
type ResumeService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (rs *ResumeService) Default() (thriftservice.ResumeService, error) {

	_err := rs.Init()
	if _err != nil {
		return nil, _err
	}

	return rs.Open()
}

// Init service get addr
func (rs *ResumeService) Init() error {

	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosQueryResume)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("resume query micro service is 0")
	}

	rs.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

// Open service open addr
func (rs *ResumeService) Open() (thriftservice.ResumeService, error) {

	transportSocket, err := thrift.NewTSocket(rs.addr)
	if err != nil {
		gatherthrift.Logger.Log("during", "thrift.NewTSocket", "err", err)
		return nil, err
	}
	rs.trans = gatherthrift.TransportFactory.GetTransport(transportSocket)
	// defer trans.Close()
	if err := rs.trans.Open(); err != nil {
		gatherthrift.Logger.Log("during", "thrift transport.Open", "err", err)
		return nil, err
	}
	cli := thriftresume.NewResumeServiceClientFactory(rs.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.ResumeService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (rs *ResumeService) Close() {
	if rs.trans.IsOpen() {
		rs.trans.Close()
	}
}
