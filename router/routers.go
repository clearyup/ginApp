package router

import (
	"ginProject/middleware"

	"github.com/gin-gonic/gin"
)

func InitPublicRouter(r *gin.Engine) {
	//用户服务路由
	InitUserRouter(r)
}
func InitProtectedRouters(r *gin.Engine) {
	// 使用jwt中间件
	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())

	//订单服务路由
	InitOrderRouter(protected)
}
