package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautaromdelgado/gin-go/internal/User/dto"
)

func (h *UserHandler) Create(c *gin.Context) {
	var user dto.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "oooops.. something went wrong"})
		return
	}
	if err := h.processor.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oooops.. something went wrong"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}
