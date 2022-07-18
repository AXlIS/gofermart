package handler

import (
	"github.com/AXlIS/gofermart/internal/service"
	"github.com/AXlIS/gofermart/pkg/auth"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
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

	debitInput struct {
		Order float32 `json:"order"`
		Sum   float32 `json:"sum"`
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

		authorized := api.Group("/", h.CheckTokenHandler())
		{
			authorized.GET("/refresh", h.GetNewAccess)
			authorized.POST("/orders", h.LoadOrder)      // Загрузка номера заказа
			authorized.GET("/orders", h.GetOrders)       // Получение списка загруженных номеров заказов
			authorized.GET("/balance", h.GetUserBalance) // Получение текущего баланса пользователя
			authorized.POST("withdraw", h.Debit)         // Запрос на списание средств
			authorized.GET("/withdrawals", h.DebitInfo)  // Получение информации о выводе средств

		}
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
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, map[string]string{"access_token": accessToken})
}

func (h *Handler) LoadOrder(c *gin.Context) {

	id := c.GetString("id")

	dataBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	orderNumber, err := strconv.Atoi(string(dataBytes))

	if err := h.service.Orders.Load(id, orderNumber); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.Status(http.StatusAccepted)
}

func (h *Handler) GetOrders(c *gin.Context) {
	id := c.GetString("id")

	orders, err := h.service.Orders.Get(id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) GetUserBalance(c *gin.Context) {
	id := c.GetString("id")

	balance, err := h.service.Users.GetBalance(id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, balance)
}

func (h *Handler) Debit(c *gin.Context) {
	var input debitInput

	id := c.GetString("id")

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Users.Debit(id, input.Sum, input.Order); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.Status(http.StatusOK)
}

func (h *Handler) DebitInfo(c *gin.Context) {
	id := c.GetString("id")

	withdrawals, err := h.service.Users.GetWithdrawalsInfo(id)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("content-type", "application/json")
	c.JSON(http.StatusOK, withdrawals)
}
