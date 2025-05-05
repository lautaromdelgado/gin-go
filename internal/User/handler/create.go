package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lautaromdelgado/gin-go/internal/User/domain"
)

func (h *UserHandler) Create(c *gin.Context) {
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "oooops.. something went wrong"})
	}
	if err := h.processor.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oooops.. something went wrong"})
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}
