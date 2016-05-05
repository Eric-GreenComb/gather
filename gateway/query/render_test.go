package query

import (
	"testing"
)

// need start micro render service localhost:39030
func TestRenderOpenService(t *testing.T) {

	var _render_service RenderService
	_thrift_service, _ := _render_service.Default()
	defer _render_service.Close()

	v := _thrift_service.RenderHello("hello", "eric")

	if v != "<html><head></head><body><h1>Hello eric</h1></body></html>" {
		t.Errorf("TestOpenService error")
	}
}
