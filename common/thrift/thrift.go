package thrift

import (
	"flag"
	"os"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"
)

var Logger log.Logger
var ProtocolFactory thrift.TProtocolFactory
var TransportFactory thrift.TTransportFactory

func init() {
	var (
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")
	)
	flag.Parse()

	Logger = log.NewLogfmtLogger(os.Stdout)
	Logger = log.NewContext(Logger).With("caller", log.DefaultCaller)

	switch *thriftProtocol {
	case "compact":
		ProtocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		ProtocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		ProtocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		ProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		Logger.Log("protocol", *thriftProtocol, "err", "invalid protocol")
		os.Exit(1)
	}

	if *thriftBufferSize > 0 {
		TransportFactory = thrift.NewTBufferedTransportFactory(*thriftBufferSize)
	} else {
		TransportFactory = thrift.NewTTransportFactory()
	}
	if *thriftFramed {
		TransportFactory = thrift.NewTFramedTransportFactory(TransportFactory)
	}
}
