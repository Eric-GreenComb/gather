package bean

import ()

type Email struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	To       string `json:"to"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	Mailtype string `json:"type"`
}
