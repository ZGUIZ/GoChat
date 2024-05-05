package router

import (
	"GoChat/com/zguiz/gochat/docs"
	"GoChat/com/zguiz/gochat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//初始界面
	r.GET("/index", service.GetIndex)
	//获取用户列表
	r.GET("/user/getUserList", service.GetUserListService)
	//用户注册
	r.POST("/user/register", service.RegisUser)
	//更新用户信息
	r.POST("/user/updateUser", service.UpdateUser)
	return r
}
