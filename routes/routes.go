package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "pong",
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		nameParam := c.Param("name")
		if nameParam == "" {
			c.String(http.StatusBadRequest, "Name query parameter is required")
			return
		}
		c.String(http.StatusOK, "Hello, %s!", nameParam)
		return
	})
}
