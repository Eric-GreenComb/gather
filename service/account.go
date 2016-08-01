package service

import (
	"encoding/json"
	"errors"
	"github.com/banerwai/global/bean"

	"github.com/banerwai/gather/gateway/command"
	"github.com/banerwai/gather/gateway/query"
	// "github.com/banerwai/global"
	// "labix.org/v2/mgo/bson"
)

type AccountService struct {
}

var _account_command_service command.AccountService
var _account_query_service query.AccountService

/**
 * command section
 */

func (self *AccountService) CreateAccount(json_account string) (v string) {
	_service, _err := _account_command_service.Default()
	if _err != nil {
		return
	}
	defer _account_command_service.Close()
	v = _service.CreateAccount(json_account)
	return
}

func (self *AccountService) CreateAccountBean(account bean.Account) (v string) {
	b, err := json.Marshal(account)
	if err != nil {
		return err.Error()
	}
	v = self.CreateAccount(string(b))
	return
}

func (self *AccountService) CreateBilling(json_billing string) (v string) {
	_service, _err := _account_command_service.Default()
	if _err != nil {
		return
	}
	defer _account_command_service.Close()
	v = _service.CreateBilling(json_billing)
	return
}

func (self *AccountService) CreateBillingBean(billing bean.Billing) (v string) {
	b, err := json.Marshal(billing)
	if err != nil {
		return err.Error()
	}
	v = self.CreateBilling(string(b))
	return
}

func (self *AccountService) DealBilling(billing_id string) (v string) {
	_service, _err := _account_command_service.Default()
	if _err != nil {
		return
	}
	defer _account_command_service.Close()
	v = _service.DealBilling(billing_id)
	return
}

func (self *AccountService) CancelBilling(billing_id string) (v string) {
	_service, _err := _account_command_service.Default()
	if _err != nil {
		return
	}
	defer _account_command_service.Close()
	v = _service.CancelBilling(billing_id)
	return
}

func (self *AccountService) GenAccount(user_id string) (v string) {
	_service, _err := _account_command_service.Default()
	if _err != nil {
		return
	}
	defer _account_command_service.Close()
	v = _service.GenAccount(user_id)
	return
}

/**
 * query section
 */

func (self *AccountService) GetAccountByUserId(user_id string) (v string) {
	_service, _err := _account_query_service.Default()
	if _err != nil {
		return
	}
	defer _account_query_service.Close()
	v = _service.GetAccountByUserId(user_id)
	return
}

func (self *AccountService) GetAccountBean(user_id string) (bean.Account, error) {
	var _obj bean.Account
	_json := self.GetAccountByUserId(user_id)
	if len(_json) == 0 {
		return _obj, errors.New("account :" + user_id + " is null")
	}

	err := json.Unmarshal([]byte(_json), &_obj)
	if err != nil {
		return _obj, errors.New("account unmarshal json error")
	}
	return _obj, nil
}

func (self *AccountService) GetBillingById(id string) (v string) {
	_service, _err := _account_query_service.Default()
	if _err != nil {
		return
	}
	defer _account_query_service.Close()
	v = _service.GetBillingById(id)
	return
}

func (self *AccountService) GetBillingBean(id string) (bean.Billing, error) {
	var _obj bean.Billing
	_json := self.GetBillingById(id)
	if len(_json) == 0 {
		return _obj, errors.New("billing :" + id + " is null")
	}

	err := json.Unmarshal([]byte(_json), &_obj)
	if err != nil {
		return _obj, errors.New("billing unmarshal json error")
	}
	return _obj, nil
}

func (self *AccountService) GetDealBillingByUserId(user_id string, timestamp int64, pagesize int64) (v string) {
	_service, _err := _account_query_service.Default()
	if _err != nil {
		return
	}
	defer _account_query_service.Close()
	v = _service.GetDealBillingByUserId(user_id, timestamp, pagesize)
	return
}

func (self *AccountService) GetDealBillingBeans(user_id string, timestamp int64, pagesize int64) ([]bean.Billing, error) {
	var _objs []bean.Billing

	_json := self.GetDealBillingByUserId(user_id, timestamp, pagesize)
	if len(_json) == 0 {
		return _objs, errors.New("billings :" + user_id + "'s deal billings is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("deal billings unmarshal json error")
	}
	return _objs, nil
}

func (self *AccountService) GetBillingByUserId(user_id string, timestamp int64, pagesize int64) (v string) {
	_service, _err := _account_query_service.Default()
	if _err != nil {
		return
	}
	defer _account_query_service.Close()
	v = _service.GetBillingByUserId(user_id, timestamp, pagesize)
	return
}

func (self *AccountService) GetBillingBeans(user_id string, timestamp int64, pagesize int64) ([]bean.Billing, error) {
	var _objs []bean.Billing

	_json := self.GetBillingByUserId(user_id, timestamp, pagesize)
	if len(_json) == 0 {
		return _objs, errors.New("billings :" + user_id + "'s billings is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("billings unmarshal json error")
	}
	return _objs, nil
}
