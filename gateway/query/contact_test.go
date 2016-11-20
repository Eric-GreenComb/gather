package query

import (
	"testing"
)

// need start micro render service localhost:39050
func TestContactDefaultService(t *testing.T) {

	var _contactService ContactService
	_thriftService, _ := _contactService.Default()

	v := _thriftService.GetContactTpl("default")

	if len(v) == 0 {
		t.Errorf("GetContactTpl error")
	}
	_contactService.Close()
}
