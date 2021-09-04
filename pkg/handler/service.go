package handler

import (
	"GoNotes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) username(c *gin.Context) {
	userInfo := h.getUserInfo(c)

	c.JSON(http.StatusOK, userInfo)
}

func (h *Handler) api(c *gin.Context) {
	c.JSON(http.StatusOK, "all good, you are authorized")
}

// getUserInfo получает инфу о пользователе по id
func (h *Handler) getUserInfo(c *gin.Context) *GoNotes.UserInfo {
	userId, err := h.getUserId(c)
	if err != nil {
		logrus.Error("error: getUserInfo: can't get cookie: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return nil
	}
	userInfo, err := h.services.User.GetUserInfo(userId)
	if err != nil {
		logrus.Error("error: getUserInfo: can't find userId: " + err.Error())
		newErrorResponse(http.StatusUnauthorized, c, "you are unauthorized")
		return nil
	}

	return userInfo
}

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	session, err := h.sessions.Get(c.Request, sessionName)
	if err != nil {
		return 0, err
	}

	return session.Values["user_id"].(int), nil
}
