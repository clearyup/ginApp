package router

import (
	"ginProject/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	// 允许跨域请求
	r.Use(middleware.Cors())
	// 初始化不需要jwt认证的路由
	InitPublicRouter(r)
	// 初始化需要jwt认证的路由
	InitProtectedRouters(r)

}

// 初始化不需要jwt认证的路由
func InitPublicRouter(r *gin.Engine) {
	//用户服务路由
	InitUserRouter(r)
}

// 初始化需要jwt认证的路由
func InitProtectedRouters(r *gin.Engine) {
	// 使用jwt中间件
	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())

	//订单服务路由
	InitOrderRouter(protected)
}
