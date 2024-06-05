package main

import (
	"ginProject/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化不需要jwt认证的路由
	router.InitPublicRouter(r)

	// r.Use(middleware.JWTMiddleware())
	// 初始化需要jwt认证的路由
	router.InitProtectedRouters(r)

	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
