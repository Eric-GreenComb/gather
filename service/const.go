package service

import (
	"flag"
	"github.com/nats-io/nats"
)

var (
	NatsUrls string
)

func init() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	flag.Parse()
	NatsUrls = *urls
}
