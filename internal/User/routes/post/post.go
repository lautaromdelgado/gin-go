package post

import (
	"github.com/gin-gonic/gin"
	"github.com/lautaromdelgado/gin-go/internal/User/handler"
)

func PostUserRoutes(r *gin.Engine, userHandler handler.IUserHandler) {
	r.POST("/create/user", userHandler.Create)
}
