package handler

import (
	"GoNotes"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode"
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

	err = passwordValidate(input.Password)
	if err != nil {
		newErrorResponse(http.StatusBadRequest, "error create user:", c, err.Error())
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

func passwordValidate(password string) error {
	var number, upper bool

	count := 0

	for _, c := range password {
		switch  {
		case unicode.IsNumber(c):
			number = true
			count++
		case unicode.IsUpper(c):
			upper = true
			count++
		case unicode.IsLetter(c) || c == ' ':
			count++
		default:
			return errors.New("password must be contains only latin symbols and numbers")
		}
		}

		if count < 7 {
			return errors.New("password must be at least 7 characters long")
		}

		if !number {
			return errors.New("password must contains at least 1 number")
		}

		if !upper {
			return errors.New("password must contains at least 1 upper case letter")
		}

		return nil
}
