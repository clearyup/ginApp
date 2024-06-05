package services

import (
	"ginProject/models"
	"ginProject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserServices struct{}

func (u *UserServices) AddUser(c *gin.Context) {
	userID := "12345"
	token, err := utils.CreateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserServices) DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete user",
	})
}

func (u *UserServices) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "update user",
	})
}

type UserReq struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func (u *UserServices) QueryUser(c *gin.Context) {

	var req UserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, gin.H{
			"message": "query user param null !",
			"data":    req.Name,
		})
		return
	} else {
		userDb := &models.User{}
		user, err := userDb.GetUserByName(req.Name)
		if err != nil {
			c.JSON(200, gin.H{
				"message": "query user err from database",
				"data":    req,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "query user success",
			"data":    user,
		})

	}

}
