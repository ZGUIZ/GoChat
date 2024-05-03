package bean

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	Password      string
	Telephone     string
	Email         string
	ClientIp      string
	ClientPort    string
	IsLoginOut    bool
	LoginTime     string
	HeartbeatTime string
	LoginoutTime  string
	DeviceInfo    string
}

func (table *User) TableName() string {
	return "user"
}
