package test

import (
	bean "GoChat/com/zguiz/gochat/models"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TextGome() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/GoChat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		errors.New("数据库连接异常")
	}
	user := &bean.User{}
	db.AutoMigrate(user)
	user.Name = "狂徒张三"
	db.Create(user)

	result := db.First(user, 1)
	fmt.Printf("姓名：%s", result.Name())
	db.Model(&user).Update("Password", "123456")

}
