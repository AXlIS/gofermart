package handler

import (
	"fmt"
	"github.com/AXlIS/gofermart/internal/service"
	"github.com/AXlIS/gofermart/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	service        *service.Service
	tokenManager   auth.TokenManager
	accessTokenTTL time.Duration
}

type (
	userInput struct {
		Username string `json:"username" binding:"required,min=2,max=64"`
		Password string `json:"password" binding:"required,min=6,max=64"`
	}

	newAccessInput struct {
		RefreshToken string `json:"refresh_token"`
	}
)

func NewHandler(service *service.Service, tokenManager auth.TokenManager, accessTokenTTL time.Duration) *Handler {
	return &Handler{
		service:        service,
		tokenManager:   tokenManager,
		accessTokenTTL: accessTokenTTL,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/user")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
		api.GET("/refresh", h.CheckTokenHandler(), h.GetNewAccess)
	}

	return router
}

func (h *Handler) Register(c *gin.Context) {
	var input userInput

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Users.Register(input.Username, input.Password); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.Status(http.StatusCreated)
}

func (h *Handler) Login(c *gin.Context) {
	var input userInput

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.service.Users.Login(input.Username, input.Password)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, tokens)
}

func (h *Handler) GetNewAccess(c *gin.Context) {

	id := c.GetString("id")

	accessToken, err := h.service.Users.GetNewAccess(id)
	if err != nil {
		fmt.Println("2")
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, map[string]string{"access_token": accessToken})
}
