package command

import (
	"testing"
)

// need start micro render service localhost:6040
func TestTokenOpenService(t *testing.T) {

	var _token_service TokenService
	_thrift_service, _ := _token_service.Default()

	_thrift_service.NewToken_("ministor@126.com", 1)

	_token_service.Close()

	_token_service.Init()
	_thrift_service, _ = _token_service.Open()
	defer _token_service.Close()
	v2 := _thrift_service.DeleteToken("ministor@126.com", 1)
	if v2 != true {
		t.Errorf("DeleteToken error")
	}
}
