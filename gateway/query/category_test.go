package query

import (
	"strings"
	"testing"
)

// need start micro render service localhost:39010
func TestCategoryDefaultService(t *testing.T) {

	var _categoryService CategoryService
	_thriftService, _ := _categoryService.Default()

	v := _thriftService.Ping()

	if v != "pong" {
		t.Errorf("Ping error")
	}
	_categoryService.Close()

	_categoryService.Init()
	_thriftService, _ = _categoryService.Open()
	defer _categoryService.Close()

	_b := _thriftService.LoadCategory("category.json")
	fmt.Println(_b)
	v1 := _thriftService.GetSubCategories(10100)
	if !strings.Contains(v1, "10101") {
		t.Errorf("TestCategoryDefaultService error")
	}
}
