package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitApplication() {
	err := initProprites()
	//检查是否异常
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	initDatabase()
}

/**
 *读取配置文件
 */
func initProprites() error {
	//配置文件路径
	viper.AddConfigPath("./com/zguiz/gochat/config")
	//配置文件名
	viper.SetConfigName("application")
	//配置文件后缀
	viper.SetConfigType("yml")
	err := viper.ReadInConfig()
	return err
}

func initDatabase() {
	fmt.Println("application :", viper.Get("applicationName"))
	mysqlConfig := viper.Get("mysql")
	fmt.Println("application :", mysqlConfig)
	if mysqlConfig != nil {
		config := interfaceToMapByReflection(mysqlConfig)
		initMySql(config)
	}
}
