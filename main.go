package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lautaromdelgado/gin-go/internal/User/handler"
	"github.com/lautaromdelgado/gin-go/internal/User/processor"
	"github.com/lautaromdelgado/gin-go/internal/User/routes"
)

func main() {
	r := gin.Default()

	userProcessor := processor.NewUserProcessor()
	userHandler := handler.NewUserHandler(userProcessor)
	routes.InitUserRoutes(r, userHandler)

	r.Run(":8080")
}
