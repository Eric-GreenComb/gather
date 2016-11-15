package service

import (
	"github.com/banerwai/gather/gateway/query"
)

var _renderQueryService query.RenderService

// RenderService RenderService
type RenderService struct {
}

/**
 * query section
 */

// RenderTpl render a tpl from etcd
func (rs *RenderService) RenderTpl(tplname string, keyMmap map[string]string) (v string) {
	_service, _err := _renderQueryService.Default()
	if _err != nil {
		return
	}
	defer _renderQueryService.Close()
	v = _service.RenderTpl(tplname, keyMmap)
	return
}
