package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestAccountService(t *testing.T) {

	var _accountService AccountService
	_thriftService, _ := _accountService.Default()

	_v := _thriftService.Ping()

	if _v != "pong" {
		t.Errorf("ContactService error")
	}
	_accountService.Close()

}
