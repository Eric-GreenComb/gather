package service

import (
	"encoding/json"
	"errors"
	"github.com/banerwai/global/bean"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
	"github.com/banerwai/global"

	"labix.org/v2/mgo/bson"
)

type ContactService struct {
}

var _contact_command_service command.ContactService
var _contact_query_service query.ContactService

/**
 * command section
 */

func (self *ContactService) CreateContact(json_contact string) (v string) {
	_service, _err := _contact_command_service.Default()
	if _err != nil {
		return
	}
	defer _contact_command_service.Close()
	v = _service.CreateContact(json_contact)
	return
}

func (self *ContactService) CreateContactBean(contact bean.Contact) (v string) {
	b, err := json.Marshal(contact)
	if err != nil {
		return err.Error()
	}
	v = self.CreateContact(string(b))
	return
}

func (self *ContactService) UpdateContact(contact_id string, mmap map[string]string) (v string) {
	_service, _err := _contact_command_service.Default()
	if _err != nil {
		return
	}
	defer _contact_command_service.Close()
	v = _service.UpdateContact(contact_id, mmap)
	return
}

func (self *ContactService) ClientSignContact(contact_id string, status bool) (v string) {
	_service, _err := _contact_command_service.Default()
	if _err != nil {
		return
	}
	defer _contact_command_service.Close()
	v = _service.ClientSignContact(contact_id, status)
	return
}

func (self *ContactService) FreelancerSignContact(contact_id string, status bool) (v string) {
	_service, _err := _contact_command_service.Default()
	if _err != nil {
		return
	}
	defer _contact_command_service.Close()
	v = _service.FreelancerSignContact(contact_id, status)
	return
}

func (self *ContactService) DealContact(contact_id string, status bool) (v string) {
	_service, _err := _contact_command_service.Default()
	if _err != nil {
		return
	}
	defer _contact_command_service.Close()
	v = _service.DealContact(contact_id, status)
	return
}

/**
 * query section
 */

func (self *ContactService) GetContactTpl(tplname string) (v string) {
	_service, _err := _contact_query_service.Default()
	if _err != nil {
		return
	}
	defer _contact_query_service.Close()
	v = _service.GetContactTpl(tplname)
	return
}

func (self *ContactService) GetContact(contactid string) (v string) {
	_service, _err := _contact_query_service.Default()
	if _err != nil {
		return
	}
	defer _contact_query_service.Close()
	v = _service.GetContact(contactid)
	return
}

func (self *ContactService) GetContactBean(contactid string) (bean.Contact, error) {
	var _obj bean.Contact
	_json := self.GetContact(contactid)
	if len(_json) == 0 {
		return _obj, errors.New("contact :" + contactid + " is null")
	}

	err := json.Unmarshal([]byte(_json), &_obj)
	if err != nil {
		return _obj, errors.New("contact unmarshal json error")
	}
	return _obj, nil
}

func (self *ContactService) GetContactSignStatus(contactid string) (v string) {
	_service, _err := _contact_query_service.Default()
	if _err != nil {
		return
	}
	defer _contact_query_service.Close()
	v = _service.GetContactSignStatus(contactid)
	return
}

func (self *ContactService) GetContactSignStatusEnum(contactid string) (v int) {
	_bson := self.GetContactSignStatus(contactid)
	if len(_bson) == 0 {
		return global.ContactSign_Null
	}
	mmap := bson.M{}
	bson.Unmarshal([]byte(_bson), mmap)

	if mmap["client_signup"] == true && mmap["freelancer_signup"] == true {
		return global.ContactSign_Dealed
	}

	if mmap["client_signup"] == true {
		return global.ContactSign_Client
	}

	if mmap["freelancer_signup"] == true {
		return global.ContactSign_Freelancer
	}
	return global.ContactSign_Null
}

func (self *ContactService) GetClientContact(clientemail string) (v string) {
	_service, _err := _contact_query_service.Default()
	if _err != nil {
		return
	}
	defer _contact_query_service.Close()
	v = _service.GetClientContact(clientemail)
	return
}

func (self *ContactService) GetClientContactBeans(clientemail string) ([]bean.Contact, error) {
	var _objs []bean.Contact

	_json := self.GetClientContact(clientemail)
	if len(_json) == 0 {
		return _objs, errors.New("contacts :" + clientemail + "'s contact is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("contacts unmarshal json error")
	}
	return _objs, nil
}

func (self *ContactService) GetFreelancerContact(freelanceremail string) (v string) {
	_service, _err := _contact_query_service.Default()
	if _err != nil {
		return
	}
	defer _contact_query_service.Close()
	v = _service.GetFreelancerContact(freelanceremail)
	return
}

func (self *ContactService) GetFreelancerContactBeans(freelanceremail string) ([]bean.Contact, error) {
	var _objs []bean.Contact

	_json := self.GetFreelancerContact(freelanceremail)
	if len(_json) == 0 {
		return _objs, errors.New("contacts :" + freelanceremail + "'s contact is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("contacts unmarshal json error")
	}
	return _objs, nil
}
