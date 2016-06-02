package query

import (
	"strings"
	"testing"
)

// need start micro render service localhost:36020
func TestTokenOpenService(t *testing.T) {

	var _auth_service AuthService
	_thrift_service, _ := _auth_service.Default()
	defer _auth_service.Close()

	v := _thrift_service.Login("ministor@126.com", "a11111")
	if strings.Contains(v, "error:") {
		t.Errorf(v)
	}
}
