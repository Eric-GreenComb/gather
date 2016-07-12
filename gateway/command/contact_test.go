package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestContactService(t *testing.T) {

	var _contact_service ContactService
	_thrift_service, _ := _contact_service.Default()

	_v := _thrift_service.Ping()

	if _v != "pong" {
		t.Errorf("ContactService error")
	}
	_contact_service.Close()

}
