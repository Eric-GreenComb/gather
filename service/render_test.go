package service

import (
	"testing"
)

func TestAuth(t *testing.T) {

	var _service RenderService
	v := _service.RenderHello("hello", "eric")
	if v != "<html><head></head><body><h1>Hello eric</h1></body></html>" {
		t.Errorf("TestOpenService error")
	}

}
