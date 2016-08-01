package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestAccountService(t *testing.T) {

	var _account_service AccountService
	_thrift_service, _ := _account_service.Default()

	_v := _thrift_service.Ping()

	if _v != "pong" {
		t.Errorf("ContactService error")
	}
	_account_service.Close()

}
