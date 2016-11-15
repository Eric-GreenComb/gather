package service

import (
	"testing"
)

func TestRenderTpl(t *testing.T) {

	var _service RenderService

	_mapParse := make(map[string]string)
	_mapParse["Hi"] = "Hello"
	_mapParse["Name"] = "Eric"

	v := _service.RenderTpl("hello", _mapParse)
	if v != "<html><head></head><body><h1>Hello Eric</h1></body></html>" {
		t.Errorf("TestRenderTpl error")
	}

}
