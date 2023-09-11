package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LogRequest(c *gin.Context) {
	log.Printf("Received request: %s %s", c.Request.Method, c.Request.URL)
	c.Next()
}
