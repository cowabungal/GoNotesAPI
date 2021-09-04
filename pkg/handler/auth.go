package handler

import (
	"GoNotes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"unicode"
)

const sessionName = "session"

// signUp производит регистрацию пользователя
func (h *Handler) signUp(c *gin.Context) {
	var input GoNotes.User

	err := c.BindJSON(&input)
	if err != nil {
		logrus.Error("error get signing-up data: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	err = passwordValidate(input.Password)
	if err != nil {
		logrus.Error("error: signUp: passwordValidate: ", err.Error())
		newErrorResponse(http.StatusBadRequest, c, err.Error())
		return
	}

	err = h.services.Authorization.CreateUser(input)
	if err != nil {
		logrus.Error("error: signUp: CreateUser: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	err = h.createSession(&input, c)
	if err != nil {
		logrus.Error("error: signUp: createSession: " + err.Error())
		newErrorResponse(http.StatusUnauthorized, c, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, "user was successfully registered")
}

// signIn производит авторизацию пользователя
func (h *Handler) signIn(c *gin.Context) {
	var input GoNotes.User

	err := c.BindJSON(&input)
	if err != nil {
		logrus.Error("error: signIn: get signing-in data: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	err = h.createSession(&input, c)
	if err != nil {
		logrus.Error("error: signIn: createSession: " + err.Error())
		newErrorResponse(http.StatusUnauthorized, c, "username or password is not correct")
		return
	}

	c.JSON(http.StatusOK, "user was successfully login")
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
