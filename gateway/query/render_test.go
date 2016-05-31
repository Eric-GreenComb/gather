package query

import (
	"testing"
)

// need start micro render service localhost:39030
func TestRenderOpenService(t *testing.T) {

	var _render_service RenderService
	_thrift_service, _ := _render_service.Default()
	defer _render_service.Close()

	_map_parse := make(map[string]string)
	_map_parse["Hi"] = "Hello"
	_map_parse["Name"] = "Eric"

	v := _thrift_service.RenderTpl("hello", _map_parse)

	if v != "<html><head></head><body><h1>Hello Eric</h1></body></html>" {
		t.Errorf("RenderTpl error")
	}
}
