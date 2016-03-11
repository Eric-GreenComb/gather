package service

import (
	"testing"
)

// need start micro render service localhost:6003
func TestRenderOpenService(t *testing.T) {

	var _render RenderService
	_render_service, _ := _render.DefaultService()
	defer _render.CloseService()

	v := _render_service.RenderHello("hello", "eric")

	if v != "<html><head></head><body><h1>Hello eric</h1></body></html>" {
		t.Errorf("TestOpenService error")
	}
}
