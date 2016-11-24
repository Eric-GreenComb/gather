package service

import (
	"fmt"
	"github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
	"gopkg.in/mgo.v2/bson"
	"time"

	"testing"
)

func TestAccountService(t *testing.T) {

	var _service AccountService

	_userID := "5707cb10ae6faa1d1071a189"

	var _obj bean.Billing
	_obj.ID = bson.ObjectIdHex(_userID)
	_obj.UserID = bson.ObjectIdHex(_userID)
	_obj.PayUserID = bson.ObjectIdHex(_userID)
	_obj.ServiceID = bson.ObjectIdHex(_userID)
	_obj.LinkID = bson.ObjectIdHex(_userID)

	_obj.Operate = 1
	_obj.Currency = constant.CurrencyCNY
	_obj.Amount = 4000
	_obj.PayType = constant.PayTypeBankRemittance

	_billingID := _service.CreateBillingBean(_obj)

	if !bson.IsObjectIdHex(_billingID) {
		t.Errorf("CreateBillingBean error")
	}

	_ok := _service.DealBilling(_billingID)
	if _ok != "OK" {
		t.Errorf("DealBilling error")
	}

	_ok = _service.GenAccount(_userID)
	if _ok != "OK" {
		t.Errorf("GenAccount error")
	}

	_account, _ := _service.GetAccountBean(_userID)
	if _account.UserID.Hex() != _userID {
		t.Errorf("GetAccountBean error")
	}

	_billing, _ := _service.GetBillingBean(_billingID)
	if _billing.Amount != 4000 {
		t.Errorf("GetBillingBean error")
	}

	_billings, _ := _service.GetBillingBeans(_userID, time.Now().Unix(), 10)
	for _, _billing := range _billings {
		fmt.Println(_billing.Amount)
	}
}
