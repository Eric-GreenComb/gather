package service

import (
	"testing"

	"github.com/banerwai/micros/common/etcd"
)

// need start micro render service localhost:6003
func TestTokenOpenService(t *testing.T) {

	_addr, _ := etcd.GetValue("/banerwai/micros/token/addr")

	if _addr != "localhost:6004" {
		t.Errorf("etcd.GetValue error")
	}

	var _token TokenService
	_token_service := _token.OpenService(_addr)
	defer _token.CloseService()

	v := _token_service.VerifyToken("ministor@126.com", 1)

	if v != 1 {
		t.Errorf("TestOpenService error")
	}

	v = _token_service.VerifyToken("ministor@gmail.com", 1)

	if v != -1 {
		t.Errorf("TestOpenService error")
	}
}
