package router

import (
	services "ginProject/services"

	"github.com/gin-gonic/gin"
)

func InitOrderRouter(r *gin.RouterGroup) {

	order := r.Group("/order")
	// order.Use(middleware.JWTMiddleware())
	orderServices := &services.OrderServices{}
	{
		order.POST("/add", orderServices.AddOrder)
		order.POST("/delete", orderServices.DeleteOrder)
		order.POST("/update", orderServices.UpdateOrder)
		order.POST("/query", orderServices.QueryOrder)
	}

}
