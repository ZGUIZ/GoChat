package service

import (
	"GoChat/com/zguiz/gochat/dao"
	"github.com/gin-gonic/gin"
)

// GetUserListService
// @Tags 获取所有用户
// @Success 200 {string} json{"code":"users"}
// @Router /user/getUserList [get]
func GetUserListService(c *gin.Context) {
	list := dao.GetUserList()
	c.JSON(200, gin.H{
		"users": list,
	})
}
