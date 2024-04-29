package bean

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string
	Password      string
	Telephone     string
	Email         string
	ClientIp      string
	ClientPort    string
	IsLoginOut    bool
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64
	DeviceInfo    string
}

func (table *User) TableName() string {
	return "user"
}
