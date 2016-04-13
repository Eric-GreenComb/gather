package service

import (
	"testing"
)

// need start micro render service localhost:9010
func TestCategoryDefaultService(t *testing.T) {

	var _category_service CategoryService
	_thrift_service, _ := _category_service.Default()

	v := _thrift_service.SayHi("eric")

	if v != "hi,eric" {
		t.Errorf("TestCategoryDefaultService error")
	}
	_category_service.Close()

	_thrift_service, _ = _category_service.Open()
	defer _category_service.Close()

	v1 := _thrift_service.SayHi("ministor")
	if v1 != "hi,ministor" {
		t.Errorf("TestCategoryDefaultService error")
	}
}
