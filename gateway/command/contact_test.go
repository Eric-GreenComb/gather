package command

import (
	"testing"
)

// need start micro render service localhost:36020
func TestContactService(t *testing.T) {

	var _contactService ContactService
	_thriftService, _ := _contactService.Default()

	_v := _thriftService.Ping()

	if _v != "pong" {
		t.Errorf("ContactService error")
	}
	_contactService.Close()

}
