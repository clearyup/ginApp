package main

import (
	"ginProject/dao"
	"ginProject/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化路由
	router.InitRouters(r)
	// 初始化数据库连接 mysql
	_ = dao.GetDb()
	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
