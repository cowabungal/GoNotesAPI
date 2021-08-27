package handler

import (
	GoNotes "GoNotes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createSession создает сессию
func (h *Handler) createSession(user *GoNotes.User, c *gin.Context) error {
	user, err := h.services.CreateSession(user)
	if err != nil {
		return err
	}

	session, err := h.sessions.Get(c.Request, sessionName)
	if err != nil {
		return err
	}

	session.Values["user_id"] = user.Id

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		return err
	}

	return nil
}

// checkSession проверяет сессию(чекает на наличие и валидность куки)
func (h *Handler) checkSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := h.sessions.Get(c.Request, sessionName)
		if err != nil {
			newErrorResponse(http.StatusInternalServerError, "error: CheckSession: Can't get cookie:", c, err.Error())
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			newErrorResponse(http.StatusUnauthorized, "error: CheckSession: Can't get user_id:", c, "no user_id")
			return
		}

		err = h.services.Authorization.CheckSession(id.(int))
		if err != nil {
			newErrorResponse(http.StatusUnauthorized, "error: CheckSession: Not correct user_id:", c, err.Error())
			return
		}

		c.Set(sessionName, session)
		c.Next()
	}
}

// clearSession очищает сессию(выход из аккаунта)
func (h *Handler) clearSession(c *gin.Context) {
	session, err := h.sessions.Get(c.Request, sessionName)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: ClearSession: saving session:", c, err.Error())
		return
	}

	session.Values = nil

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: ClearSession: saving session:", c, err.Error())
		return
	}

	c.JSON(http.StatusOK, "message: user was successfully logout")
}
