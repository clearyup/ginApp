package collect

import (
	"github.com/gin-gonic/gin"
)

type CollectServices struct{}

func (t *CollectServices) AddCollect(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "add collect",
	})
}
