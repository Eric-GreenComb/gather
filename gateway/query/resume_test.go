package query

import (
	"fmt"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"testing"
)

func TestResumeDefaultService(t *testing.T) {

	var _resume_service ResumeService
	_thrift_service, _ := _resume_service.Default()

	v := _thrift_service.GetResume("5707cb10ae6faa1d1071a189")
	var resume bean.Resume
	bson.Unmarshal([]byte(v), &resume)
	fmt.Println(resume)

	if resume.AuthEmail != "ministor@126.com" {
		t.Errorf("GetResume error")
	}
	_resume_service.Close()

}
