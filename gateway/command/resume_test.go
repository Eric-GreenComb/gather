package command

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"testing"
)

// need start micro render service localhost:36020
func TestResumeService(t *testing.T) {

	var _service ResumeService
	_thrift_service, _ := _service.Default()

	var _obj bean.Resume
	_obj.Id = bson.ObjectIdHex("5707cb10ae6faa1d1071a189")
	_obj.AuthEmail = "ministor@126.com"
	_obj.UserID = bson.ObjectIdHex("5707cb10ae6faa1d1071a189")

	_obj.Phone = "12345678901"

	var lsToolandArchs []bean.ToolandArch

	var _tool1 bean.ToolandArch
	_tool1.ToolLevel = 5
	_tool1.ToolTitle = "Java"
	lsToolandArchs = append(lsToolandArchs, _tool1)

	var _tool2 bean.ToolandArch
	_tool2.ToolLevel = 2
	_tool2.ToolTitle = "Go"
	lsToolandArchs = append(lsToolandArchs, _tool2)

	_obj.ToolandArchs = lsToolandArchs

	b, _ := json.Marshal(_obj)
	v := _thrift_service.AddResume(string(b))
	if v != "OK" {
		t.Errorf("AddResume error")
	}

	_map_update := make(map[string]string)
	_map_update["auth_email"] = "ministor@126.com"
	_map_update["phone"] = "13811111112"
	userid := "5707cb10ae6faa1d1071a189"
	v = _thrift_service.UpdateResumeBase(userid, _map_update)

	if v != "OK" {
		t.Errorf("UpdateResume error")
	}
	_service.Close()

}
