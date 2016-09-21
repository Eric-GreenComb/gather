package flagparse

import (
	"flag"
	"github.com/nats-io/nats"
)

var (
	// NatsUrls nats url
	NatsUrls string
	// ThriftProtocol thrift protocol
	ThriftProtocol string
	// ThriftBufferSize thrift buffer size
	ThriftBufferSize int
	// ThriftFramed thrift framed
	ThriftFramed bool
	// BanerwaiWebPort banerwai web port
	BanerwaiWebPort string
	// BanerwaiAPIPort banerwai api port
	BanerwaiAPIPort string
	// BanerwaiAPIKey banerwai api secret key
	BanerwaiAPIKey string
	// BanerwaiAPICheckSign if check sign 2 banerwai api
	BanerwaiAPICheckSign bool
)

func init() {

	var (
		natsUrls         = flag.String("nats.urls", nats.DefaultURL, "The nats server URLs (separated by comma)")
		banerwaiWebPort  = flag.String("banerwai.web.port", ":3000", "the web listen port")
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")

		banerwaiAPIPort      = flag.String("banerwai.api.port", ":3000", "the api listen port")
		banerwaiAPIKey       = flag.String("banerwai.api.key", "private_key", "key for api")
		banerwaiAPICheckSign = flag.Bool("banerwai.api.check.sign", false, "if check api sign")
	)

	flag.Parse()
	NatsUrls = *natsUrls

	BanerwaiWebPort = *banerwaiWebPort

	ThriftProtocol = *thriftProtocol
	ThriftBufferSize = *thriftBufferSize
	ThriftFramed = *thriftFramed

	BanerwaiAPIPort = *banerwaiAPIPort
	BanerwaiAPIKey = *banerwaiAPIKey
	BanerwaiAPICheckSign = *banerwaiAPICheckSign
}
