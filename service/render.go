package service

import (
	"github.com/banerwai/gather/gateway/query"
)

var _query_service query.RenderService

type RenderService struct {
}

/**
 * query section
 */

func (self *RenderService) RenderHello(tmpl, name string) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.RenderHello(tmpl, name)
	return
}
