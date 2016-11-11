package service

import (
	"github.com/banerwai/gather/gateway/query"

	"github.com/banerwai/global/bean"

	"encoding/json"
)

// CategoryService CategoryService
type CategoryService struct {
}

var _categoryQueryService query.CategoryService

/**
 * query section
 */

// LoadCategory load category json 2 mem
func (cs *CategoryService) LoadCategory(path string) (v bool) {
	_service, _err := _categoryQueryService.Default()
	if _err != nil {
		return
	}
	defer _categoryQueryService.Close()
	v = _service.LoadCategory(path)
	return
}

// GetCategories get categories json
func (cs *CategoryService) GetCategories() (v string) {
	_service, _err := _categoryQueryService.Default()
	if _err != nil {
		return
	}
	defer _categoryQueryService.Close()
	v = _service.GetCategories()
	return
}

// GetCategoriesBean get category beans
func (cs *CategoryService) GetCategoriesBean() []bean.Category {
	var cats []bean.Category
	_json := cs.GetCategories()
	if len(_json) == 0 {
		return cats
	}
	err := json.Unmarshal([]byte(_json), &cats)
	if err != nil {
		return nil
	}
	return cats
}

// GetSubCategories get subcategories json
func (cs *CategoryService) GetSubCategories(serialnumber int32) (v string) {
	_service, _err := _categoryQueryService.Default()
	if _err != nil {
		return
	}
	defer _categoryQueryService.Close()
	v = _service.GetSubCategories(serialnumber)
	return
}

// GetSubCategoriesBean get subcategory beans
func (cs *CategoryService) GetSubCategoriesBean(serialnumber int32) []bean.SubCategory {
	var subs []bean.SubCategory
	_json := cs.GetSubCategories(serialnumber)
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
