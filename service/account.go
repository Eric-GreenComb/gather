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

// AccountService AccountService
type AccountService struct {
}

var _accountCommandService command.AccountService
var _accountQueryService query.AccountService

/**
 * command section
 */

// CreateAccount create account by account json
func (as *AccountService) CreateAccount(jsonAccount string) (v string) {
	_service, _err := _accountCommandService.Default()
	if _err != nil {
		return
	}
	defer _accountCommandService.Close()
	v = _service.CreateAccount(jsonAccount)
	return
}

// CreateAccountBean create account by account bean
func (as *AccountService) CreateAccountBean(account bean.Account) (v string) {
	b, err := json.Marshal(account)
	if err != nil {
		return err.Error()
	}
	v = as.CreateAccount(string(b))
	return
}

// CreateBilling create billing by billing json
func (as *AccountService) CreateBilling(jsonBilling string) (v string) {
	_service, _err := _accountCommandService.Default()
	if _err != nil {
		return
	}
	defer _accountCommandService.Close()
	v = _service.CreateBilling(jsonBilling)
	return
}

// CreateBillingBean create billing by billing bean
func (as *AccountService) CreateBillingBean(billing bean.Billing) (v string) {
	b, err := json.Marshal(billing)
	if err != nil {
		return err.Error()
	}
	v = as.CreateBilling(string(b))
	return
}

// DealBilling deal a billing
func (as *AccountService) DealBilling(billingID string) (v string) {
	_service, _err := _accountCommandService.Default()
	if _err != nil {
		return
	}
	defer _accountCommandService.Close()
	v = _service.DealBilling(billingID)
	return
}

// CancelBilling cancel a billing
func (as *AccountService) CancelBilling(billingID string) (v string) {
	_service, _err := _accountCommandService.Default()
	if _err != nil {
		return
	}
	defer _accountCommandService.Close()
	v = _service.CancelBilling(billingID)
	return
}

// GenAccount gen account by billing
func (as *AccountService) GenAccount(userID string) (v string) {
	_service, _err := _accountCommandService.Default()
	if _err != nil {
		return
	}
	defer _accountCommandService.Close()
	v = _service.GenAccount(userID)
	return
}

/**
 * query section
 */

// GetAccountByUserID get account json by userID
func (as *AccountService) GetAccountByUserID(userID string) (v string) {
	_service, _err := _accountQueryService.Default()
	if _err != nil {
		return
	}
	defer _accountQueryService.Close()
	v = _service.GetAccountByUserID(userID)
	return
}

// GetAccountBean get account bean by userID
func (as *AccountService) GetAccountBean(userID string) (bean.Account, error) {
	var _obj bean.Account
	_json := as.GetAccountByUserID(userID)
	if len(_json) == 0 {
		return _obj, errors.New("account :" + userID + " is null")
	}

	err := json.Unmarshal([]byte(_json), &_obj)
	if err != nil {
		return _obj, errors.New("account unmarshal json error")
	}
	return _obj, nil
}

// GetBillingByID get billing json by ID
func (as *AccountService) GetBillingByID(ID string) (v string) {
	_service, _err := _accountQueryService.Default()
	if _err != nil {
		return
	}
	defer _accountQueryService.Close()
	v = _service.GetBillingByID(ID)
	return
}

// GetBillingBean get billing bean by ID
func (as *AccountService) GetBillingBean(ID string) (bean.Billing, error) {
	var _obj bean.Billing
	_json := as.GetBillingByID(ID)
	if len(_json) == 0 {
		return _obj, errors.New("billing :" + ID + " is null")
	}

	err := json.Unmarshal([]byte(_json), &_obj)
	if err != nil {
		return _obj, errors.New("billing unmarshal json error")
	}
	return _obj, nil
}

// GetDealBillingByUserID get dealed billings json by userID
func (as *AccountService) GetDealBillingByUserID(userID string, timestamp int64, pagesize int64) (v string) {
	_service, _err := _accountQueryService.Default()
	if _err != nil {
		return
	}
	defer _accountQueryService.Close()
	v = _service.GetDealBillingByUserID(userID, timestamp, pagesize)
	return
}

// GetDealBillingBeans get dealed billing beans by userID
func (as *AccountService) GetDealBillingBeans(userID string, timestamp int64, pagesize int64) ([]bean.Billing, error) {
	var _objs []bean.Billing

	_json := as.GetDealBillingByUserID(userID, timestamp, pagesize)
	if len(_json) == 0 {
		return _objs, errors.New("billings :" + userID + "'s deal billings is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("deal billings unmarshal json error")
	}
	return _objs, nil
}

// GetBillingByUserID get billings json by userID
func (as *AccountService) GetBillingByUserID(userID string, timestamp int64, pagesize int64) (v string) {
	_service, _err := _accountQueryService.Default()
	if _err != nil {
		return
	}
	defer _accountQueryService.Close()
	v = _service.GetBillingByUserID(userID, timestamp, pagesize)
	return
}

// GetBillingBeans get billing beans by userID
func (as *AccountService) GetBillingBeans(userID string, timestamp int64, pagesize int64) ([]bean.Billing, error) {
	var _objs []bean.Billing

	_json := as.GetBillingByUserID(userID, timestamp, pagesize)
	if len(_json) == 0 {
		return _objs, errors.New("billings :" + userID + "'s billings is null")
	}

	err := json.Unmarshal([]byte(_json), &_objs)
	if err != nil {
		return _objs, errors.New("billings unmarshal json error")
	}
	return _objs, nil
}
