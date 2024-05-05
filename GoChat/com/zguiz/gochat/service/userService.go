package service

import (
	"GoChat/com/zguiz/gochat/dao"
	bean "GoChat/com/zguiz/gochat/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		c.JSON(200, gin.H{
			"messageCode": MESSAGE_CODE_ERROR,
			"message":     err.Error(),
		})
		return
	}

	dao.RegisterUser(param)
	c.JSON(200, gin.H{
		"messageCode": MESSAGE_CODE_SUCCESS,
		"message":     "账户注册成功!",
	})
}

// UpdateUser
// @Summary 修改用户信息
// @Tags 用户模块
// @param name query string false "用户名"
// @param oldPassword query string false "旧密码"
// @param password query string false "密码"
// @param confirmPassword query string false "确认密码"
// @param email query string false "邮件"
// @param telephone query string false "电话号码"
// @Success 200 {string} json{"code":"users"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	param := bean.User{}
	param.Name = c.Query("name")
	param.Password = c.Query("oldPassword")

	userList := dao.FindUser(param)
	var user bean.User
	var messageCode = MESSAGE_CODE_ERROR
	var message string
	if userList != nil && len(userList) > 0 {
		//直接取第一个
		user = userList[0]
		user.Password = c.Query("password")
		user.Email = c.Query("email")
		user.Telephone = c.Query("telephone")

		validate := validator.New()
		err := validate.Struct(user)
		if err != nil {
			messageCode = MESSAGE_CODE_ERROR
			message = err.Error()
			c.JSON(200, gin.H{
				"messageCode": messageCode,
				"message":     message,
			})
			return
		}

		user.Password = c.Query("password")
		confirmPassword := c.Query("confirmPassword")
		//检查验证密码
		if confirmPassword != user.Password {
			messageCode = MESSAGE_CODE_ERROR
			message = "密码与确认密码不一致!"
		} else {
			//修改账户信息
			dao.UpdateUser(user)
			messageCode = MESSAGE_CODE_SUCCESS
			message = "账户信息修改成功!"
		}
	} else {
		messageCode = MESSAGE_CODE_ERROR
		message = "账户信息不存在!"
	}

	c.JSON(200, gin.H{
		"messageCode": messageCode,
		"message":     message,
	})
}
