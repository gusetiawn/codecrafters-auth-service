package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
	"io"
	"net/http"
)

func (h *handler) Login(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "failed read request",
		})
		return
	}

	var request models.Login
	if err = json.Unmarshal(jsonData, &request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "failed decode request",
		})
		return
	}

	result, err := h.service.Login(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	if result == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Unauthorized User Login",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Success",
		"token":   result,
	})
	return
}
