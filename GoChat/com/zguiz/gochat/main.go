package main

import (
	"GoChat/com/zguiz/gochat/router"
	"GoChat/com/zguiz/gochat/utils"
)

func main() {
	//初始化配置
	utils.InitApplication()
	//启动服务器
	r := router.Router()
	r.Run(":8081")
}
