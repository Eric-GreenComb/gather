package query

import (
	"github.com/apache/thrift/lib/go/thrift"

	thriftclient "github.com/banerwai/micros/query/render/client/thrift"
	thriftservice "github.com/banerwai/micros/query/render/service"
	thriftrender "github.com/banerwai/micros/query/render/thrift/gen-go/render"

	"errors"
	gatherthrift "github.com/banerwai/gather/common/thrift"
	"github.com/banerwai/global/constant"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	"github.com/banerwai/gommon/etcd"
)

// RenderService RenderService
type RenderService struct {
	trans thrift.TTransport
	addr  string
}

// Default service init and open
func (rs *RenderService) Default() (thriftservice.RenderService, error) {
	_err := rs.Init()
	if _err != nil {
		return nil, _err
	}

	return rs.Open()
}

// Init service get addr
func (rs *RenderService) Init() error {
	_addrs, _err := etcd.GetServicesByName(constant.EtcdKeyMicrosQueryRender)

	if _err != nil {
		return _err
	}
	if len(_addrs) == 0 {
		return errors.New("render query micro service is 0")
	}

	rs.addr = _addrs[banerwaicrypto.GetRandomItNum(len(_addrs))]

	return nil
}

// Open service open addr
func (rs *RenderService) Open() (thriftservice.RenderService, error) {

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
	cli := thriftrender.NewRenderServiceClientFactory(rs.trans, gatherthrift.ProtocolFactory)

	var svc thriftservice.RenderService
	svc = thriftclient.New(cli, gatherthrift.Logger)

	return svc, err
}

// Close service close
func (rs *RenderService) Close() {
	if rs.trans.IsOpen() {
		rs.trans.Close()
	}
}
