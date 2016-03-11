package user

import (
	"testing"
)

// need start micro render service localhost:6003
func TestFindPwd(t *testing.T) {

	var _password PasswordService
	_password.FindPwd("ministor@126.com")

}
