package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lautaromdelgado/gin-go/internal/User/handler"
	"github.com/lautaromdelgado/gin-go/internal/User/routes/post"
)

func InitUserRoutes(r *gin.Engine, userHandler handler.IUserHandler) {
	post.PostUserRoutes(r, userHandler)
}
