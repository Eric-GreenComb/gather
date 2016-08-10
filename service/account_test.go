package service

import (
	"github.com/banerwai/global"
	"github.com/banerwai/global/bean"
	"labix.org/v2/mgo/bson"
	"time"

	"testing"
)

func TestAccountService(t *testing.T) {

	var _service AccountService

	_user_id := "5707cb10ae6faa1d1071a189"

	var _obj bean.Billing
	_obj.Id = bson.ObjectIdHex(_user_id)
	_obj.UserId = bson.ObjectIdHex(_user_id)
	_obj.ProfileId = bson.ObjectIdHex(_user_id)

	_obj.Operate = 1
	_obj.Currency = global.CURRENCY_CNY
	_obj.Amount = 4000
	_obj.PayType = global.PayType_BankRemittance

	_billing_id := _service.CreateBillingBean(_obj)

	if !bson.IsObjectIdHex(_billing_id) {
		t.Errorf("CreateBillingBean error")
	}

	_ok := _service.DealBilling(_billing_id)
	if _ok != "OK" {
		t.Errorf("DealBilling error")
	}

	_ok = _service.GenAccount(_user_id)
	if _ok != "OK" {
		t.Errorf("GenAccount error")
	}

	_account, _ := _service.GetAccountBean(_user_id)
	if _account.UserId.Hex() != _user_id {
		t.Errorf("GetAccountBean error")
	}

	_billing, _ := _service.GetBillingBean(_billing_id)
	if _billing.Amount != 4000 {
		t.Errorf("GetBillingBean error")
	}

	_billings, _ := _service.GetBillingBeans(_user_id, time.Now().Unix(), 10)
	if len(_billings) < 3 {
		t.Errorf("GetBillingBeans error")
	}
}
