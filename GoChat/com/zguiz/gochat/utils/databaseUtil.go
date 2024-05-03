package utils

import (
	bean "GoChat/com/zguiz/gochat/models"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

func initMySql(config map[string]interface{}) {
	//获取连接串
	url := getLinkUrl(config)
	//数据库连接
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		errors.New("数据库连接异常:" + err.Error())
		//logrus.Error("数据库连接异常:" + err.Error())
	}
	user := &bean.User{}
	user.Name = "狂徒张三"
	db.Create(user)
	result := db.First(user, 1)
	fmt.Printf("姓名：%s", result.Name())
}

// 组织连接串
func getLinkUrl(config map[string]interface{}) string {
	//ip
	ip := config["ip"]
	//端口
	port := config["port"]
	//用户名
	user := config["user"]
	//密码
	password := config["password"]
	//访问数据库
	database := config["database"]
	//组装连接串
	//"root:123456@tcp(127.0.0.1:3306)/GoChat?charset=utf8mb4&parseTime=True&loc=Local"
	builder := strings.Builder{}
	builder.WriteString(interfaceToString(user))
	builder.WriteString(":")
	builder.WriteString(interfaceToString(password))
	builder.WriteString("@tcp(")
	builder.WriteString(interfaceToString(ip))
	builder.WriteString(":")
	builder.WriteString(interfaceToString(port))
	builder.WriteString(")/")
	builder.WriteString(interfaceToString(database))
	builder.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	linkUrl := builder.String()
	//logrus.Debug("数据库连接串：" + linkUrl)
	return linkUrl
}
