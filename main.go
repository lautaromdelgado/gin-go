package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lautaromdelgado/gin-go/routes"
)

func main() {
	r := gin.Default()

	routes.InitializeRoutes(r)

	r.Run(":8080")
}
