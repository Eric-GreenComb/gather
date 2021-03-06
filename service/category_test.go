package service

import (
	"strings"
	"testing"
)

func TestAuth(t *testing.T) {

	var _service CategoryService
	b := _service.LoadCategory("category")
	if b != true {
		t.Errorf("LoadCategory error")
	}

	_json := _service.GetSubCategories(10100)
	if !strings.Contains(_json, "10101") {
		t.Errorf("GetSubCategories error")
	}

	v := _service.GetSubCategoriesBean(10200)
	if len(v) != 10 {
		t.Errorf("GetSubCategories error")
	}

}
