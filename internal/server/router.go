package server

import (
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, handlers *ServerHandlers) {
	r.GET("/ping", handlers.Ping)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.POST("/users", handlers.GetUser)
}
