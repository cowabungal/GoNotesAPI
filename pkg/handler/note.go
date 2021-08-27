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
		newErrorResponse(http.StatusInternalServerError, "error: getNotes: Can't get cookie:", c, err.Error())
	}

	var note GoNotes.Note

	err = c.BindJSON(&note)
	if err != nil {
		newErrorResponse(http.StatusBadRequest, "error: getNotes: Can't get note data:", c, err.Error())
		return
	}

	noteId, err := h.services.Note.Add(userId, &note)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: getNotes: Can't add notes:", c, err.Error())
	}

	c.JSON(http.StatusOK, noteId)
}

func (h *Handler) updateNotes(c *gin.Context) {

}

func (h *Handler) deleteNotes(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: getNotes: Can't get cookie:", c, err.Error())
	}

	var noteId GoNotes.Note

	err = c.BindJSON(&noteId)
	if err != nil {
		newErrorResponse(http.StatusBadRequest, "error: getNotes: Сan't get note data:", c, err.Error())
		return
	}

	err = h.services.Note.Delete(noteId.Id, userId)
	if err != nil {
		newErrorResponse(http.StatusInternalServerError, "error: getNotes: Can't add notes:", c, err.Error())
	}

	c.JSON(http.StatusOK, "message: note was successfully deleted")
}
