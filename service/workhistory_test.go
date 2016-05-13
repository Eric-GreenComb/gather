package service

import (
	"testing"
)

func TestWorkHistoryBase(t *testing.T) {

	var _service WorkHistoryService
	v := _service.AddWorkHistory("email")

	if v != "email" {
		t.Errorf("AddWorkHistory error")
	}

	v = _service.UpdateWorkHistory("123456")

	if v != "123456" {
		t.Errorf("UpdateWorkHistory error")
	}

	v = _service.GetWorkHistory("1")

	if v != "1" {
		t.Errorf("GetWorkHistory error")
	}
}
