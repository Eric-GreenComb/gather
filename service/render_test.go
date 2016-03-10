package service

import (
	"testing"

	"github.com/banerwai/micros/common/etcd"
)

// need start micro render service localhost:6003
func TestRenderOpenService(t *testing.T) {

	_addr, _ := etcd.GetValue("/banerwai/micros/render/addr")

	if _addr != "localhost:6003" {
		t.Errorf("etcd.GetValue error")
	}

	var _render RenderService
	_render_service := _render.OpenService(_addr)
	defer _render.CloseService()

	v := _render_service.RenderHello("hello", "eric")

	if v != "<html><head></head><body><h1>Hello eric</h1></body></html>" {
		t.Errorf("TestOpenService error")
	}
}
