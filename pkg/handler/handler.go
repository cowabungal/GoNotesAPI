package handler

import (
	"GoNotes/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// Handler хранит данные из сервисов (service)
type Handler struct {
	services *service.Service
	sessions *sessions.CookieStore
}

// NewHandler создает новый объект *Handler
func NewHandler(service *service.Service, session *sessions.CookieStore) *Handler {
	return &Handler{services: service, sessions: session}
}

// InitRoutes инициализирует routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.NoRoute(h.notFound)

	router.POST("/auth/sign-up", h.signUp)
	router.POST("/auth/sign-in", h.signIn)

	authRoutes := router.Group("/api")
	authRoutes.Use(h.checkSession())
	{
		authRoutes.GET("", h.api)
		authRoutes.GET("/logout", h.clearSession)
		userRoutes := authRoutes.Group("/user")
		{
			userRoutes.GET("/username", h.username)

			userRoutes.GET("/notes", h.getNotes)
			userRoutes.POST("/notes", h.addNotes)
			userRoutes.PUT("/notes", h.updateNotes)
			userRoutes.DELETE("/notes", h.deleteNotes)

		}
	}

	return router
}
