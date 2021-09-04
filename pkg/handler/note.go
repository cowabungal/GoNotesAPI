package handler

import (
	"GoNotes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		logrus.Error("error: getNotes: can't get cookie: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	var notes []*GoNotes.Note

	notes, err = h.services.Note.GetAll(userId)
	if err != nil {
		logrus.Error("error: getNotes: can't find notes: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	if notes == nil {
		logrus.Error("error: getNotes: user has no notes")
		c.JSON(http.StatusOK, response{"user has no notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (h *Handler) addNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		logrus.Error("error: addNotes: can't get cookie: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	var note GoNotes.Note

	err = c.BindJSON(&note)
	if err != nil {
		logrus.Error("error: addNotes: can't get note data: " + err.Error())
		newErrorResponse(http.StatusBadRequest, c, "error reading note data")
		return
	}

	noteId, err := h.services.Note.Add(userId, &note)
	if err != nil {
		logrus.Error("error: addNotes: can't add notes : " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, noteId)
}

func (h *Handler) updateNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		logrus.Error("error: updateNotes: can't get cookie: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	var note GoNotes.Note

	err = c.BindJSON(&note)
	if err != nil {
		logrus.Error("error: updateNotes: can't get note data: " + err.Error())
		newErrorResponse(http.StatusBadRequest, c, "error reading note data")
		return
	}

	noteId, err := h.services.Note.Update(userId, &note)
	if err != nil {
		logrus.Error("error: updateNotes: can't update notes: " + err.Error())
		newErrorResponse(http.StatusBadRequest, c, "user has no note with this id or bad data")
		return
	}

	c.JSON(http.StatusOK, noteId)
}

func (h *Handler) deleteNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		logrus.Error("error: deleteNotes: can't get cookie: " + err.Error())
		newErrorResponse(http.StatusInternalServerError, c, "something went wrong")
		return
	}

	var note GoNotes.Note

	err = c.BindJSON(&note)
	if err != nil {
		logrus.Error("error: deleteNotes: can't get note data: " + err.Error())
		newErrorResponse(http.StatusBadRequest, c, "error reading note data")
		return
	}

	err = h.services.Note.Delete(note.Id, userId)
	if err != nil {
		logrus.Error("error: deleteNotes: can't delete notes: " + err.Error())
		newErrorResponse(http.StatusBadRequest, c, "user has no note with this id")
		return
	}

	c.JSON(http.StatusOK, response{"note was successfully deleted"})
}
