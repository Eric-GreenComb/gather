package query

import (
	"testing"
)

// need start micro render service localhost:39050
func TestContactDefaultService(t *testing.T) {

	var _contact_service ContactService
	_thrift_service, _ := _contact_service.Default()

	v := _thrift_service.GetContactTpl("default")

	if len(v) == 0 {
		t.Errorf("GetContactTpl error")
	}
	_contact_service.Close()
}
