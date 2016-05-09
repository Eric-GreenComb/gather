package service

import (
	"github.com/banerwai/gather/gateway/query"

	"github.com/banerwai/gather/dto"

	"encoding/json"
)

var _query_service query.CategoryService

type CategoryService struct {
}

/**
 * query section
 */

func (self *CategoryService) LoadCategory(path string) (v bool) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.LoadCategory(path)
	return
}

func (self *CategoryService) GetCategories() (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetCategories()
	return
}

func (self *CategoryService) GetDtoCategories() []dto.Category {
	var cats []dto.Category
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
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetSubCategories(serialnumber)
	return
}

func (self *CategoryService) GetDtoSubCategories(serialnumber int32) []dto.SubCategory {
	var subs []dto.SubCategory
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
