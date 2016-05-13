package service

import (
	"github.com/banerwai/global/bean"
	"testing"
)

func TestTokenADV(t *testing.T) {

	var _service TokenService
	_service.NewToken_("ministor@126.com", bean.TokenActiveEmail)

	_v := _service.VerifyToken("ministor@126.com", int64(bean.TokenActiveEmail), _service.GetOverHours(bean.TokenActiveEmail))
	if _v != 1 {
		t.Errorf("VerifyToken error")
	}

	v := _service.DeleteToken("ministor@126.com", int64(bean.TokenActiveEmail))
	if v != true {
		t.Errorf("DeleteToken error")
	}
}
