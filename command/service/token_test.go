package service

import (
	"testing"

	"fmt"
)

// need start micro render service localhost:6040
func TestTokenOpenService(t *testing.T) {

	var _token_service TokenService
	_thrift_service, _ := _token_service.Default()

	v := _thrift_service.NewToken_("ministor@126.com", 1)
	fmt.Println(v)

	v1 := _thrift_service.VerifyToken("ministor@126.com", 1)
	if v1 != 1 {
		t.Errorf("VerifyToken error")
	}
	_token_service.Close()

	_thrift_service, _ = _token_service.Open()
	defer _token_service.Close()
	v2 := _thrift_service.DeleteToken("ministor@126.com", 1)
	if v2 != true {
		t.Errorf("DeleteToken error")
	}
}
