package handler

import (
	"GoNotes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

// checkSession проверяет сессию (чекает на наличие и валидность куки)
func (h *Handler) checkSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := h.sessions.Get(c.Request, sessionName)
		if err != nil {
			logrus.Error("error: checkSession: can't get cookies: " + err.Error())
			newErrorResponse(http.StatusInternalServerError, c, "can't get cookies")
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			logrus.Error("error: checkSession: can't get user_id: no user_id ")
			newErrorResponse(http.StatusUnauthorized, c, "you are unauthorized")
			return
		}

		err = h.services.Authorization.CheckSession(id.(int))
		if err != nil {
			logrus.Error("error: CheckSession: Not correct user_id: " + err.Error())
			newErrorResponse(http.StatusUnauthorized, c, "you are unauthorized")
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
		logrus.Error("error: ClearSession: get session: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	session.Values = nil

	err = session.Save(c.Request, c.Writer)
	if err != nil {
		logrus.Error("error: ClearSession: saving session: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, "user was successfully logout")
}
