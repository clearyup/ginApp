package router

import (
	services "ginProject/services"
	"ginProject/services/collect"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {

	//user services
	userServices := &services.UserServices{}
	user := r.Group("/user")
	{
		user.POST("/add", userServices.AddUser)
		user.POST("/delete", userServices.DeleteUser)
		user.POST("/update", userServices.UpdateUser)
		user.POST("/query", userServices.QueryUser)
	}

	CollectServices := &collect.CollectServices{}
	collect := r.Group("/collect")

	{
		collect.POST("/add", CollectServices.AddCollect)
	}
}
