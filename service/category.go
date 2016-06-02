package service

import (
	"github.com/banerwai/gather/gateway/query"

	"github.com/banerwai/global/bean"

	"encoding/json"
)

var _category_query_service query.CategoryService

type CategoryService struct {
}

/**
 * query section
 */

func (self *CategoryService) LoadCategory(path string) (v bool) {
	_service, _err := _category_query_service.Default()
	if _err != nil {
		return
	}
	defer _category_query_service.Close()
	v = _service.LoadCategory(path)
	return
}

func (self *CategoryService) GetCategories() (v string) {
	_service, _err := _category_query_service.Default()
	if _err != nil {
		return
	}
	defer _category_query_service.Close()
	v = _service.GetCategories()
	return
}

func (self *CategoryService) GetCategoriesBean() []bean.Category {
	var cats []bean.Category
	_json := self.GetCategories()
	if len(_json) == 0 {
		return cats
	}
	err := json.Unmarshal([]byte(_json), &cats)
	if err != nil {
		return nil
	}
	return cats
}

func (self *CategoryService) GetSubCategories(serialnumber int32) (v string) {
	_service, _err := _category_query_service.Default()
	if _err != nil {
		return
	}
	defer _category_query_service.Close()
	v = _service.GetSubCategories(serialnumber)
	return
}

func (self *CategoryService) GetSubCategoriesBean(serialnumber int32) []bean.SubCategory {
	var subs []bean.SubCategory
	_json := self.GetSubCategories(serialnumber)
	if len(_json) == 0 {
		return subs
	}
	err := json.Unmarshal([]byte(_json), &subs)
	if err != nil {
		return nil
	}
	return subs
}

/**
 * extend section
 */
