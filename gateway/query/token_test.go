package query

import (
	"testing"

	"github.com/banerwai/gather/command/bean"
)

// need start micro render service localhost:6020
func TestTokenOpenService(t *testing.T) {

	var _token_service TokenService
	_thrift_service, _ := _token_service.Default()
	defer _token_service.Close()

	_v := _thrift_service.VerifyToken("ministor@126.com", bean.TokenActiveEmail, getOverHours(bean.TokenActiveEmail))
	if _v != 1 {
		t.Errorf("VerifyToken error")
	}
}

func getOverHours(ttype int64) float64 {
	switch ttype {

	// 0 - 2.0
	case bean.TokenPwd:
		return bean.PwdOverHours

	// 1 - 48.0
	case bean.TokenActiveEmail:
		return bean.ActiveEmailOverHours

	// 2 - 2.0
	case bean.TokenUpdateEmail:
		return bean.UpdateEmailOverHours
	}
	return 0
}
