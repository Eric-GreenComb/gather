package command

import (
	"testing"
)

// need start micro render service localhost:36040
func TestTokenOpenService(t *testing.T) {

	var _tokenService TokenService
	_thriftService, _ := _tokenService.Default()

	_thriftService.NewToken_("ministor@126.com", 1)

	_tokenService.Close()

	_tokenService.Init()
	_thriftService, _ = _tokenService.Open()
	defer _tokenService.Close()
	v2 := _thriftService.DeleteToken("ministor@126.com", 1)
	if v2 != true {
		t.Errorf("DeleteToken error")
	}
}
