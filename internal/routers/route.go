package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	c "main.go/internal/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong"})
	})

	item := r.Group("/item")
	{
		item.POST("", c.Create().CreateItem)
	}
	return r
}
