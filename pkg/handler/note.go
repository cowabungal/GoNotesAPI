package handler

import (
	GoNotes "GoNotes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: getNotes: Can't get cookie:", c, err.Error())
	}

	var notes []*GoNotes.Note

	notes, err = h.services.Note.GetAll(userId)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: getNotes: Can't find notes:", c, err.Error())
	}

	c.JSON(http.StatusOK, notes)
}

func (h *Handler) addNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: addNotes: Can't get cookie: ", c, err.Error())
	}

	var note GoNotes.Note

	err = c.BindJSON(&note)
	if err != nil {
		newErrorResponse(http.StatusBadRequest, "error: addNotes: Can't get note data: ", c, err.Error())
		return
	}

	noteId, err := h.services.Note.Add(userId, &note)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: addNotes: Can't add notes: ", c, err.Error())
	}

	c.JSON(http.StatusOK, noteId)
}

func (h *Handler) updateNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: updateNotes: Can't get cookie: ", c, err.Error())
	}

	var note GoNotes.Note

	err = c.BindJSON(&note)
	if err != nil {
		newErrorResponse(http.StatusBadRequest, "error: updateNotes: Сan't get note data: ", c, err.Error())
		return
	}

	noteId, err := h.services.Note.Update(userId, &note)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: updateNotes: Can't update notes: ", c, err.Error())
		return
	}

	c.JSON(http.StatusOK, noteId)
}

func (h *Handler) deleteNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: deleteNotes: Can't get cookie: ", c, err.Error())
	}

	var note GoNotes.Note

	err = c.BindJSON(&note)
	if err != nil {
		newErrorResponse(http.StatusBadRequest, "error: deleteNotes: Сan't get note data: ", c, err.Error())
		return
	}

	err = h.services.Note.Delete(note.Id, userId)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: deleteNotes: Can't add notes: ", c, err.Error())
	}

	c.JSON(http.StatusOK, "message: note was successfully deleted")
}
