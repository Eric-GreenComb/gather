package gearman

import (
	"log"

	"github.com/banerwai/micros/common/etcd"
	"github.com/mikespook/gearman-go/client"
)

var GearmanAddr string

// var GearmanClient *client.Client
var GearmanResponseHandler client.ResponseHandler

func init() {
	var _err error
	GearmanAddr, _err = etcd.GetValue("/banerwai/gearman/conn")
	if _err != nil {
		return
	}

	GearmanResponseHandler = func(resp *client.Response) {
		switch resp.DataType {
		case client.WorkException:
			fallthrough
		case client.WorkFail:
			fallthrough
		case client.WorkComplate:
			if data, err := resp.Result(); err == nil {
				log.Printf("RESULT: %v\n", data)
			} else {
				log.Printf("RESULT: %s\n", err)
			}
		case client.WorkWarning:
			fallthrough
		case client.WorkData:
			if data, err := resp.Update(); err == nil {
				log.Printf("UPDATE: %v\n", data)
			} else {
				log.Printf("UPDATE: %v, %s\n", data, err)
			}
		case client.WorkStatus:
			if data, err := resp.Status(); err == nil {
				log.Printf("STATUS: %v\n", data)
			} else {
				log.Printf("STATUS: %s\n", err)
			}
		default:
			log.Printf("UNKNOWN: %v", resp.Data)
		}
	}
}
