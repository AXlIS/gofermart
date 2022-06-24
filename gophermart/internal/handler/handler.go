package handler

import (
	"github.com/AXlIS/gofermart/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service

}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/user")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
	}

	return router
}

func (h *Handler) Register(c *gin.Context) {

}

func (h *Handler) Login(c *gin.Context) {

}

