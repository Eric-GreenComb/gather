package service

import (
	"testing"
)

// need start micro render service localhost:6030
func TestCategoryDefaultService(t *testing.T) {

	var _category CategoryService
	_category_service, _ := _category.DefaultService()
	defer _category.CloseService()

	v := _category_service.SayHi("eric")

	if v != "hi,eric" {
		t.Errorf("TestCategoryDefaultService error")
	}
}
