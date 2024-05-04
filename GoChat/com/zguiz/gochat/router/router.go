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
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserListService)
	r.POST("/user/register", service.RegisUser)
	return r
}
