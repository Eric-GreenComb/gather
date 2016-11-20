package query

import (
	"testing"

	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
)

// need start micro render service localhost:6020
func TestUserOpenService(t *testing.T) {

	var _userService UserService
	_thriftService, _ := _userService.Default()
	defer _userService.Close()

	_v := _thriftService.GetUser("ministor@126.com")
	_user := bean.UserDto{}
	bson.Unmarshal([]byte(_v), &_user)
	if _user.Email != "ministor@126.com" {
		t.Errorf("GetUser error")
	}
}
