package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) notFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, nil)
}

func newErrorResponse(statusCode int, c *gin.Context, message string) {
	c.AbortWithStatusJSON(statusCode, response{message})
}
