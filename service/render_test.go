package service

import (
	"testing"
)

func TestRenderTpl(t *testing.T) {

	var _service RenderService

	_map_parse := make(map[string]string)
	_map_parse["Hi"] = "Hello"
	_map_parse["Name"] = "Eric"

	v := _service.RenderTpl("hello", _map_parse)
	if v != "<html><head></head><body><h1>Hello Eric</h1></body></html>" {
		t.Errorf("TestRenderTpl error")
	}

}
