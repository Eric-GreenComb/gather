package bean

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`          // 必须要设置bson:"_id" 不然mgo不会认为是主键
	Invited      bson.ObjectId `bson:"invited" json:"invited"` // invited by someone
	Email        string        `bson:"email" json:"email"`     // 全是小写
	Pwd          string        `bson:"pwd" json:"-"`
	Actived      bool          `bson:"actived" json:"actived"`
	CreatedTime  time.Time     `bson:"createdtime" json:"createdtime"`
	LastActivity time.Time     `bson:"lastactivity" json:"lastactivity"`
}
