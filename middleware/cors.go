package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckCors() gin.HandlerFunc {
	//这里可以处理一些别的逻辑
	return func(c *gin.Context) {
		// 定义一个origin的map，只有在字典中的key才允许跨域请求
		var allowOrigins = map[string]struct{}{
			"http://127.0.0.1:8888":       {},
			"https://www.yangyanxing.com": {},
		}
		origin := c.Request.Header.Get("Origin") //请求头部
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		if origin != "" {
			if _, ok := allowOrigins[origin]; ok {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
				c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
				c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}
		c.Next()
	}
}
