package thrift

import (
	"os"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/banerwai/gather/common/flagparse"
	"github.com/go-kit/kit/log"
)

var (
	// Logger logger
	Logger log.Logger
	// ProtocolFactory thrift protocol factory
	ProtocolFactory thrift.TProtocolFactory
	// TransportFactory thrift transport factory
	TransportFactory thrift.TTransportFactory
)

func init() {

	Logger = log.NewLogfmtLogger(os.Stdout)
	Logger = log.NewContext(Logger).With("caller", log.DefaultCaller)

	switch flagparse.ThriftProtocol {
	case "compact":
		ProtocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		ProtocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		ProtocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		ProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		Logger.Log("protocol", flagparse.ThriftProtocol, "err", "invalid protocol")
		os.Exit(1)
	}

	if flagparse.ThriftBufferSize > 0 {
		TransportFactory = thrift.NewTBufferedTransportFactory(flagparse.ThriftBufferSize)
	} else {
		TransportFactory = thrift.NewTTransportFactory()
	}
	if flagparse.ThriftFramed {
		TransportFactory = thrift.NewTFramedTransportFactory(TransportFactory)
	}
}
