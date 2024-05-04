package service

import (
	"GoChat/com/zguiz/gochat/dao"
	bean "GoChat/com/zguiz/gochat/models"
	"github.com/gin-gonic/gin"
)

// GetUserListService
// @Summary 查询所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code":"users"}
// @Router /user/getUserList [get]
func GetUserListService(c *gin.Context) {
	list := dao.GetUserList()
	c.JSON(200, gin.H{
		"users": list,
	})
}

// RegisUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param confirmPassword query string false "确认密码"
// @param email query string false "邮件"
// @param telephone query string false "电话号码"
// @Success 200 {string} json{"code":"users"}
// @Router /user/register [post]
func RegisUser(c *gin.Context) {
	param := bean.User{}
	param.Name = c.Query("name")
	param.Password = c.Query("password")
	param.Email = c.Query("email")
	param.Telephone = c.Query("telephone")
	param.IsLoginOut = false
	confirmPassword := c.Query("confirmPassword")
	if confirmPassword != param.Password {
		c.JSON(200, gin.H{
			"messageCode": MESSAGE_CODE_ERROR,
			"message":     "密码与确认密码不一致!",
		})
		return
	}
	dao.RegisterUser(param)
	c.JSON(200, gin.H{
		"messageCode": MESSAGE_CODE_SUCCESS,
		"message":     "账户注册成功!",
	})
}
