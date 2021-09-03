package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func (h *Handler) notFound(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, nil)
}

func newErrorResponse(statusCode int, info string, c *gin.Context, message string) {
	logrus.Error(info + message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
