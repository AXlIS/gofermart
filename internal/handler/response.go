package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Error struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	log.Error().Msg("error: " + message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
