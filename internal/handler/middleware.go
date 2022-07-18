package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = `Authorization`
)

func (h *Handler) CheckTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := h.parseAuthorization(c)

		if err != nil {
			errorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		c.Set("id", id)
		c.Next()
	}
}

func (h *Handler) parseAuthorization(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return h.tokenManager.Parse(headerParts[1])
}
