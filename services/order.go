package services

import "github.com/gin-gonic/gin"

type OrderServices struct{}

func (o *OrderServices) AddOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add order",
	})
}

func (o *OrderServices) DeleteOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete order",
	})
}

func (o *OrderServices) UpdateOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update order",
	})
}

func (o *OrderServices) QueryOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "query order",
	})
}
