package service

import (
	"github.com/banerwai/gather/gateway/query"
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

func (self *CategoryService) GetSubCategories(serialnumber int32) (v string) {
	_service, _ := _query_service.Default()
	defer _query_service.Close()
	v = _service.GetSubCategories(serialnumber)
	return
}
