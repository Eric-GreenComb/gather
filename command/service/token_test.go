package service

import (
	"testing"
)

// need start micro render service localhost:6003
func TestTokenOpenService(t *testing.T) {

	var _token TokenService
	_token_service, _ := _token.DefaultService()
	defer _token.CloseService()

	_token_service.NewToken_("ministor@126.com", 1)
	v := _token_service.VerifyToken("ministor@126.com", 1)

	if v != 1 {
		t.Errorf("TestOpenService error")
	}
}
