package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestResumeService(t *testing.T) {

	var _service ResumeService
	_thrift_service, _ := _service.Default()

	v := _thrift_service.AddResume("email")

	if v != "email" {
		t.Errorf("AddResume error")
	}

	v = _thrift_service.UpdateResume("123456")

	if v != "123456" {
		t.Errorf("UpdateResume error")
	}
	_service.Close()

}
