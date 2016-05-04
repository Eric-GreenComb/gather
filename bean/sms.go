package bean

import ()

// alidayu sms struct
type SMS struct {
	RecNum          string `json:"rec_num"`
	SmsFreeSignName string `json:"sms_free_sign_name"`
	SmsTemplateCode string `json:"sms_template_code"`
	SmsParam        string `json:"sms_param"`
}
