package handler

import (
	"GoNotes"
	"github.com/gin-gonic/gin"
	"net/http"
)

const sessionName = "session"

// signUp производит регистрацию пользователя
func (h *Handler) signUp(c *gin.Context) {
	var input GoNotes.User

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error get signing-up data", c, err.Error())
		return
	}

	err = h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error create user:", c, err.Error())
		return
	}

	err = h.createSession(&input, c)
	if err != nil {
		newErrorResponse(http.StatusUnauthorized, "error create session:", c, err.Error())
		return
	}

	c.JSON(http.StatusOK, "message: user was successfully registered")
}

// signIn производит авторизацию пользователя
func (h *Handler) signIn(c *gin.Context) {
	var input GoNotes.User

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error get signing-in data", c, err.Error())
		return
	}

	err = h.createSession(&input, c)
	if err != nil {
		newErrorResponse(http.StatusUnauthorized, "error signing-in:", c, err.Error())
		return
	}

	c.JSON(http.StatusOK, "message: user was successfully login")
}
