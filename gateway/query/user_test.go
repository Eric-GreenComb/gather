package query

import (
	"testing"

	"github.com/banerwai/gather/bean"
	"labix.org/v2/mgo/bson"
)

// need start micro render service localhost:6020
func TestUserOpenService(t *testing.T) {

	var _user_service UserService
	_thrift_service, _ := _user_service.Default()
	defer _user_service.Close()

	_v := _thrift_service.GetUser("ministor@126.com")
	_user := bean.UserDto{}
	bson.Unmarshal([]byte(_v), &_user)
	if _user.Email != "ministor@126.com" {
		t.Errorf("GetUser error")
	}
}
