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

func FindUser(user bean.User) []bean.User {
	param := &bean.User{Name: user.Name, Password: user.Password}
	userList := []bean.User{}
	utils.DatabaseConn.Find(&userList, param)
	return userList
}

func UpdateUser(user bean.User) {
	utils.DatabaseConn.Model(&user).Updates(bean.User{Name: user.Name, Password: user.Password, Telephone: user.Telephone, Email: user.Email})
}
