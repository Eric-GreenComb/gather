package query

import (
	"testing"
)

// need start micro render service localhost:39050
func TestAccountDefaultService(t *testing.T) {

	var _account_service AccountService
	_thrift_service, _ := _account_service.Default()

	v := _thrift_service.Ping()

	if v != "pong" {
		t.Errorf("Ping error")
	}
	_account_service.Close()
}
