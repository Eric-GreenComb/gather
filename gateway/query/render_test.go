package query

import (
	"testing"
)

// need start micro render service localhost:39030
func TestRenderOpenService(t *testing.T) {

	var _renderService RenderService
	_thriftService, _ := _renderService.Default()
	defer _renderService.Close()

	_mapParse := make(map[string]string)
	_mapParse["Hi"] = "Hello"
	_mapParse["Name"] = "Eric"

	v := _thriftService.RenderTpl("hello", _mapParse)

	if v != "<html><head></head><body><h1>Hello Eric</h1></body></html>" {
		t.Errorf("RenderTpl error")
	}
}
