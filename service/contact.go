package service

import (
	"encoding/json"
	"errors"
	"github.com/banerwai/global/bean"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
	"github.com/banerwai/global/constant"

	"labix.org/v2/mgo/bson"
)

// ContactService ContactService
type ContactService struct {
}

var _contactCommandService command.ContactService
var _contactQueryService query.ContactService

/**
 * command section
 */

// CreateContact create contact by json
func (cs *ContactService) CreateContact(jsonContact string) (v string) {
	_service, _err := _contactCommandService.Default()
	if _err != nil {
		return
	}
	defer _contactCommandService.Close()
	v = _service.CreateContact(jsonContact)
	return
}

// CreateContactBean create contact by contact bean
func (cs *ContactService) CreateContactBean(contact bean.Contact) (v string) {
	b, err := json.Marshal(contact)
	if err != nil {
		return err.Error()
	}
	v = cs.CreateContact(string(b))
	return
}

// UpdateContact update contact by map
func (cs *ContactService) UpdateContact(contactID string, mmap map[string]string) (v string) {
	_service, _err := _contactCommandService.Default()
	if _err != nil {
		return
	}
	defer _contactCommandService.Close()
	v = _service.UpdateContact(contactID, mmap)
	return
}

// ClientSignContact client sign contact
func (cs *ContactService) ClientSignContact(contactID string, status bool) (v string) {
	_service, _err := _contactCommandService.Default()
	if _err != nil {
		return
	}
	defer _contactCommandService.Close()
	v = _service.ClientSignContact(contactID, status)
	return
}

// FreelancerSignContact freelancer sign contact
func (cs *ContactService) FreelancerSignContact(contactID string, status bool) (v string) {
	_service, _err := _contactCommandService.Default()
	if _err != nil {
		return
	}
	defer _contactCommandService.Close()
	v = _service.FreelancerSignContact(contactID, status)
	return
}

// DealContact client/freelancer sign contact and deal contact
func (cs *ContactService) DealContact(contactID string, status bool) (v string) {
	_service, _err := _contactCommandService.Default()
	if _err != nil {
		return
	}
	defer _contactCommandService.Close()
	v = _service.DealContact(contactID, status)
	return
}

/**
 * query section
 */

// GetContactTpl get contact by tpl
func (cs *ContactService) GetContactTpl(tplname string) (v string) {
	_service, _err := _contactQueryService.Default()
	if _err != nil {
		return
	}
	defer _contactQueryService.Close()
	v = _service.GetContactTpl(tplname)
	return
}

// GetContact get contact json by ID
func (cs *ContactService) GetContact(contactID string) (v string) {
	_service, _err := _contactQueryService.Default()
	if _err != nil {
		return
	}
	defer _contactQueryService.Close()
	v = _service.GetContact(contactID)
	return
}

// GetContactBean get contact bean by ID
func (cs *ContactService) GetContactBean(contactID string) (bean.Contact, error) {
	var _obj bean.Contact
	_json := cs.GetContact(contactID)
	if len(_json) == 0 {
		return _obj, errors.New("contact :" + contactID + " is null")
	}

	err := json.Unmarshal([]byte(_json), &_obj)
	if err != nil {
		return _obj, errors.New("contact unmarshal json error")
	}
	return _obj, nil
}

// GetContactSignStatus get contact sign status
func (cs *ContactService) GetContactSignStatus(contactID string) (v string) {
	_service, _err := _contactQueryService.Default()
	if _err != nil {
		return
	}
	defer _contactQueryService.Close()
	v = _service.GetContactSignStatus(contactID)
	return
}

// GetContactSignStatusEnum get contact sign enum
func (cs *ContactService) GetContactSignStatusEnum(contactID string) (v int) {
	_bson := cs.GetContactSignStatus(contactID)
	if len(_bson) == 0 {
		return constant.ContactSignNull
	}
	mmap := bson.M{}
	bson.Unmarshal([]byte(_bson), mmap)

	if mmap["client_signup"] == true && mmap["freelancer_signup"] == true {
		return constant.ContactSignDealed
	}

	if mmap["client_signup"] == true {
		return constant.ContactSignClient
	}

	if mmap["freelancer_signup"] == true {
		return constant.ContactSignFreelancer
	}
	return constant.ContactSignNull
}

// GetClientContact get client contacts json
func (cs *ContactService) GetClientContact(clientemail string) (v string) {
	_service, _err := _contactQueryService.Default()
	if _err != nil {
		return
	}
	defer _contactQueryService.Close()
	v = _service.GetClientContact(clientemail)
	return
}

// GetClientContactBeans get client contact beans
func (cs *ContactService) GetClientContactBeans(clientemail string) ([]bean.Contact, error) {
	var _objs []bean.Contact

	_json := cs.GetClientContact(clientemail)
	if len(_json) == 0 {
		return _objs, errors.New("contacts :" + clientemail + "'s contact is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("contacts unmarshal json error")
	}
	return _objs, nil
}

// GetFreelancerContact get freelancer contacts json
func (cs *ContactService) GetFreelancerContact(freelanceremail string) (v string) {
	_service, _err := _contactQueryService.Default()
	if _err != nil {
		return
	}
	defer _contactQueryService.Close()
	v = _service.GetFreelancerContact(freelanceremail)
	return
}

// GetFreelancerContactBeans get freelancer contact beans
func (cs *ContactService) GetFreelancerContactBeans(freelanceremail string) ([]bean.Contact, error) {
	var _objs []bean.Contact

	_json := cs.GetFreelancerContact(freelanceremail)
	if len(_json) == 0 {
		return _objs, errors.New("contacts :" + freelanceremail + "'s contact is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("contacts unmarshal json error")
	}
	return _objs, nil
}
