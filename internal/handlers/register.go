package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gusetiawn/codecrafters-auth-service/internal/models"
	"io"
	"net/http"
)

func (h *handler) Register(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "failed read request",
		})
		return
	}

	var request models.Register
	if err = json.Unmarshal(jsonData, &request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "failed decode request",
		})
		return
	}

	err = h.service.Register(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Success",
	})
	return
}
