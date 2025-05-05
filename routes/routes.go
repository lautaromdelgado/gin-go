package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Nombre   string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func InitializeRoutes(r *gin.Engine) {
	// Obtener un "Hello, World!" en la ruta raíz
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Obtener un "pong" en la ruta /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "pong",
		})
	})

	// Obtener un "Hello, {name}!" con el parametro de la ruta que le enviemos
	r.GET("/hello/:name", func(c *gin.Context) {
		nameParam := c.Param("name")
		if nameParam == "" {
			c.String(http.StatusBadRequest, "Name query parameter is required")
			return
		}
		c.String(http.StatusOK, "Hello, %s!", nameParam)
		return
	})

	// Registrar un usuario (nombre, apellido y email)
	r.POST("/usuarios", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "oooops.. something went wrong"})
			return
		}
		if user.Nombre == "" || user.Apellido == "" || user.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ups.. some fields are empty"})
		}
		c.JSON(http.StatusCreated, gin.H{
			"status": "success",
			"data":   user,
		})
	})
}
