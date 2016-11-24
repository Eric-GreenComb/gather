package service

import (
	"encoding/json"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"gopkg.in/mgo.v2/bson"
	"strings"

	"testing"
)

func TestContactService(t *testing.T) {

	var _service ContactService

	_contactTpl := _service.GetContactTpl("default")
	if !strings.Contains(_contactTpl, "{{.ContactNumber}}") {
		t.Errorf("GetContactTpl error")
	}

	var _obj bean.Contact
	_obj.ID = bson.NewObjectId()
	_obj.ClientEmail = "ministor@gmail.com"
	_obj.FreeLancerEmail = "ministor@126.com"

	_mmap := prepareContactParam()
	_bParam, _ := json.Marshal(_mmap)
	_obj.ContactTpl = _contactTpl
	_obj.TplParam = string(_bParam)

	_contactID := _service.CreateContactBean(_obj)
	if !bson.IsObjectIdHex(_contactID) {
		t.Errorf("CreateContactBean error")
	}

	_mmapUpdatePrepare := updateContactParam()
	_bParamUpdate, _ := json.Marshal(_mmapUpdatePrepare)
	_mmapUpdate := make(map[string]string)
	_mmapUpdate["tpl_param"] = string(_bParamUpdate)

	_ok := _service.UpdateContact(_contactID, _mmapUpdate)
	if _ok != "OK" {
		t.Errorf("UpdateContact error")
	}

	_ok = _service.DealContact(_contactID, true)
	if _ok != "OK" {
		t.Errorf("DealContact error")
	}

	_contact, _err := _service.GetContactBean(_contactID)
	if _err != nil {
		t.Errorf("GetContactBean error")
	}
	if !bson.IsObjectIdHex(_contact.ID.Hex()) {
		t.Errorf("GetContactBean error")
	}
	if len(_contact.ContactContent) == 0 {
		t.Errorf("GetContactBean error")
	}

	_statusEnum := _service.GetContactSignStatusEnum(_contactID)
	if constant.ContactSignDealed != _statusEnum {
		t.Errorf("GetContactBean error")
	}

	_contactClient, _ := _service.GetClientContactBeans("ministor@gmail.com")
	if len(_contactClient) != 1 {
		t.Errorf("GetClientContactBeans error")
	}

	_contactFreelancer, _ := _service.GetFreelancerContactBeans("ministor@126.com")
	if len(_contactFreelancer) != 1 {
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
