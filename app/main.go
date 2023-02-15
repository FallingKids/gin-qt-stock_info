package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/config"
	"github.com/gin-qt-business/app/router"
)

func init() {
	// 初始化配置文件
	config.ViperInit()

	// 初始化mysql连接
	config.InitMysqlDB()
}

func main() {
	r := gin.Default()

	router.Router(r)

	_ = r.Run()
}
