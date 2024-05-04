package dao

import (
	bean "GoChat/com/zguiz/gochat/models"
	"GoChat/com/zguiz/gochat/utils"
)

/**
 *查询所有用户
 */
func GetUserList() []bean.User {
	param := &bean.User{}
	userList := []bean.User{}
	utils.DatabaseConn.Find(&userList, param)
	return userList
}

/**
 * 用户注册
 */
func RegisterUser(user bean.User) {
	utils.DatabaseConn.Create(&user)
}
