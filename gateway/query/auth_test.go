package query

import (
	"strings"
	"testing"
)

// need start micro render service localhost:36020
func TestTokenOpenService(t *testing.T) {

	var _authService AuthService
	_thriftService, _ := _authService.Default()
	defer _authService.Close()

	v := _thriftService.Login("ministor@126.com", "a11111")
	if strings.Contains(v, "error:") {
		t.Errorf(v)
	}
}
