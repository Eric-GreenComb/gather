package service

import (
	"testing"
)

func TestResumeBase(t *testing.T) {

	var _service ResumeService
	v := _service.AddResume("email")

	if v != "email" {
		t.Errorf("AddResume error")
	}

	v = _service.UpdateResume("123456")

	if v != "123456" {
		t.Errorf("UpdateResume error")
	}

	v = _service.GetResume("1")

	if v != "1" {
		t.Errorf("GetResume error")
	}
}
