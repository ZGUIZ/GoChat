package dao

import (
	bean "GoChat/com/zguiz/gochat/models"
	"GoChat/com/zguiz/gochat/utils"
)

func GetUserList() []bean.User {
	param := &bean.User{}
	userList := []bean.User{}
	utils.DatabaseConn.Find(&userList, param)
	return userList
}
