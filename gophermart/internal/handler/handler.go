package handler

import (
	"github.com/AXlIS/gofermart/internal/service"
	"github.com/AXlIS/gofermart/pkg/auth"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	service        *service.Service
	tokenManager   auth.TokenManager
	accessTokenTTL time.Duration
}

func NewHandler(service *service.Service, tokenManager auth.TokenManager, accessTokenTTL time.Duration) *Handler {
	return &Handler{
		service: service,
		tokenManager: tokenManager,
		accessTokenTTL: accessTokenTTL,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/user")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
		api.GET("/refresh", h.GetNewRefresh)
	}

	return router
}

func (h *Handler) Register(c *gin.Context) {

}

func (h *Handler) Login(c *gin.Context) {

}

func (h *Handler) GetNewRefresh(c *gin.Context) {

}