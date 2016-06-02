package service

import (
	"github.com/banerwai/gather/gateway/query"
)

var _render_query_service query.RenderService

type RenderService struct {
}

/**
 * query section
 */

func (self *RenderService) RenderTpl(tplname string, key_mmap map[string]string) (v string) {
	_service, _err := _render_query_service.Default()
	if _err != nil {
		return
	}
	defer _render_query_service.Close()
	v = _service.RenderTpl(tplname, key_mmap)
	return
}
