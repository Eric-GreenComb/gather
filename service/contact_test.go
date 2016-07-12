package service

import (
	"encoding/json"
	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"strings"

	"testing"
)

func TestContactService(t *testing.T) {

	var _service ContactService

	_contact_tpl := _service.GetContactTpl("default")
	if !strings.Contains(_contact_tpl, "{{.ContactNumber}}") {
		t.Errorf("GetContactTpl error")
	}

	var _obj bean.Contact
	_obj.Id = bson.NewObjectId()
	_obj.ClientEmail = "ministor@gmail.com"
	_obj.FreeLancerEmail = "ministor@126.com"

	_mmap := prepareContactParam()
	_bParam, _ := json.Marshal(_mmap)
	_obj.ContactTpl = _contact_tpl
	_obj.TplParam = string(_bParam)

	_contact_id := _service.CreateContactBean(_obj)
	if !bson.IsObjectIdHex(_contact_id) {
		t.Errorf("CreateContactBean error")
	}

	_mmap_update_prepare := updateContactParam()
	_bParam_update, _ := json.Marshal(_mmap_update_prepare)
	_mmap_update := make(map[string]string)
	_mmap_update["tpl_param"] = string(_bParam_update)

	_ok := _service.UpdateContact(_contact_id, _mmap_update)
	if _ok != "OK" {
		t.Errorf("UpdateContact error")
	}

	_ok = _service.DealContact(_contact_id, true)
	if _ok != "OK" {
		t.Errorf("DealContact error")
	}

	_contact, _err := _service.GetContactBean(_contact_id)
	if _err != nil {
		t.Errorf("GetContactBean error")
	}
	if !bson.IsObjectIdHex(_contact.Id.Hex()) {
		t.Errorf("GetContactBean error")
	}
	if len(_contact.ContactContent) == 0 {
		t.Errorf("GetContactBean error")
	}

	_status_enum := _service.GetContactSignStatusEnum(_contact_id)
	if global.ContactSign_Dealed != _status_enum {
		t.Errorf("GetContactBean error")
	}

	_contact_client, _ := _service.GetClientContactBeans("ministor@gmail.com")
	if len(_contact_client) != 1 {
		t.Errorf("GetClientContactBeans error")
	}

	_contact_freelancer, _ := _service.GetFreelancerContactBeans("ministor@126.com")
	if len(_contact_freelancer) != 1 {
		t.Errorf("GetFreelancerContactBeans error")
	}
}

func prepareContactParam() map[string]string {
	_map := make(map[string]string)
	_map["ProjectName"] = "banerwai服务"
	_map["ClientName"] = "ministor@gmail.com"
	_map["FreeLancerName"] = "ministor@126.com"
	_map["ExpectedTime"] = "4weeks"
	_map["PayRate"] = "100Y / hour"
	_map["PaymentMethod"] = "week"
	return _map
}

func updateContactParam() map[string]string {
	_map := make(map[string]string)
	_map["ProjectName"] = "banerwai服务"
	_map["ClientName"] = "ministor@gmail.com"
	_map["FreeLancerName"] = "ministor@126.com"
	_map["ExpectedTime"] = "12weeks"
	_map["PayRate"] = "150Y / hour"
	_map["PaymentMethod"] = "week"
	return _map
}
