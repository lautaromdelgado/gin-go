package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lautaromdelgado/gin-go/internal/User/processor"
)

type IUserHandler interface {
	Create(c *gin.Context)
}

type UserHandler struct {
	processor processor.IUserProcessor
}

func NewUserHandler(processor processor.IUserProcessor) IUserHandler {
	return &UserHandler{
		processor: processor,
	}
}
