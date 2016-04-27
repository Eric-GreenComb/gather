package service

import (
	"strings"
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

	_category_service.Init()
	_thrift_service, _ = _category_service.Open()
	defer _category_service.Close()

	_b := _thrift_service.LoadCategory("category.json")
	fmt.Println(_b)
	v1 := _thrift_service.GetSubCategories(10100)
	if !strings.Contains(v1, "10101") {
		t.Errorf("TestCategoryDefaultService error")
	}
}
