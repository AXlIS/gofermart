package handler

import (
	g "github.com/AXlIS/gofermart"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

type (
	AccrualHandler struct{}
)

func NewAccrualHandler() *AccrualHandler {
	return &AccrualHandler{}
}

func (h *AccrualHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/api/orders/:order", h.OrderHandler)

	return router
}

func (h *AccrualHandler) OrderHandler(c *gin.Context) {
	order := c.Param("order")

	orderNumber, err := strconv.Atoi(order)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if rand.Intn(10)%2 == 0 {
		accrual := float64(orderNumber % 359)
		c.JSON(http.StatusOK, g.Order{
			Number:   orderNumber,
			Status:  "PROCESSED",
			Accrual: &accrual,
		})
		return
	}

	switch rand.Intn(2) {
	case 0:
		c.JSON(http.StatusOK,
			g.Order{Number: orderNumber, Status: "REGISTERED"})
	case 1:
		c.JSON(http.StatusOK,
			g.Order{Number: orderNumber, Status: "PROCESSING"})
	case 2:
		c.JSON(http.StatusOK,
			g.Order{Number: orderNumber, Status: "INVALID"})
	}
}
