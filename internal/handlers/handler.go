package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gusetiawn/codecrafters-auth-service/internal/services"
)

type Handler interface {
	Ping(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUser(c *gin.Context)
}

type handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return &handler{
		service: service,
	}
}
