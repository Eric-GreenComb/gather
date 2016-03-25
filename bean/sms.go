package bean

import ()

type Sms struct {
	Name    string `json:"name"`
	Pwd     string `json:"pwd"`
	Content string `json:"content"`
	Mobile  string `json:"mobile"`
	Sign    string `json:"sign"`
	Extno   string `json:"extno"`
}
