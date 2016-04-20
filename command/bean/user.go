package bean

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`                  // 必须要设置bson:"_id" 不然mgo不会认为是主键
	Invited      bson.ObjectId `bson:"invited" json:"invited"`         // invited by someone
	Email        string        `bson:"email" json:"email"`             // 全是小写
	Verified     bool          `bson:"verified" json:"verified"`       // Email是否已验证过?
	Username     string        `bson:"username" json:"username"`       // 不区分大小写, 全是小写
	UsernameRaw  string        `bson:"usernameraw" json:"usernameraw"` // 可能有大小写
	Pwd          string        `bson:"pwd" json:"-"`
	CreatedTime  time.Time     `bson:"createdtime" json:"createdtime"`
	LastActivity time.Time     `bson:"lastactivity" json:"lastactivity"`
}
