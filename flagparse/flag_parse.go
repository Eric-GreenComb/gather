package flagparse

import (
	"flag"
	"github.com/nats-io/nats"
)

var (
	NatsUrls         string
	BanerwaiWebPort  string
	ThriftProtocol   string
	ThriftBufferSize int
	ThriftFramed     bool

	BanerwaiApiPort string
	BanerwaiApiKey  string
)

func init() {

	var (
		natsUrls         = flag.String("nats.urls", nats.DefaultURL, "The nats server URLs (separated by comma)")
		banerwaiWebPort  = flag.String("banerwai.web.port", ":3000", "the web listen port")
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")

		banerwaiApiPort = flag.String("banerwai.api.port", ":3000", "the api listen port")
		banerwaiApiKey  = flag.String("banerwai.api.key", "private_key", "key for api")
	)

	flag.Parse()
	NatsUrls = *natsUrls

	BanerwaiWebPort = *banerwaiWebPort

	ThriftProtocol = *thriftProtocol
	ThriftBufferSize = *thriftBufferSize
	ThriftFramed = *thriftFramed

	BanerwaiApiPort = *banerwaiApiPort
	BanerwaiApiKey = *banerwaiApiKey
}
