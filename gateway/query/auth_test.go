package query

import (
	"testing"
)

// need start micro render service localhost:6020
func TestTokenOpenService(t *testing.T) {

	var _auth_service AuthService
	_thrift_service, _ := _auth_service.Default()
	defer _auth_service.Close()

	v := _thrift_service.Login("emailOrUsername", "pwd")
	if v != "emailOrUsernamepwd" {
		t.Errorf("Login error")
	}
}
