package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) notFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, nil)
}

func newErrorResponse(statusCode int, c *gin.Context, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
