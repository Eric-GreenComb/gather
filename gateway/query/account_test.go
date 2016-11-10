package query

import (
	"testing"
)

// need start micro render service localhost:39050
func TestAccountDefaultService(t *testing.T) {

	var _accountService AccountService
	_thriftSrvice, _ := _accountService.Default()

	v := _thriftSrvice.Ping()

	if v != "pong" {
		t.Errorf("Ping error")
	}
	_accountService.Close()
}
