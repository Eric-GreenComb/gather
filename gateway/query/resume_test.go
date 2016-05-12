package query

import (
	"testing"
)

func TestResumeDefaultService(t *testing.T) {

	var _resume_service ResumeService
	_thrift_service, _ := _resume_service.Default()

	v := _thrift_service.GetResume("1234567")

	if v != "1234567" {
		t.Errorf("GetResume error")
	}
	_resume_service.Close()

}
