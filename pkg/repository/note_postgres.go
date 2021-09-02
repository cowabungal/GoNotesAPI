package repository

import (
	"GoNotes"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type NoteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) GetAll(userId int) ([]*GoNotes.Note, error) {
	query := fmt.Sprintf("SELECT * from %s WHERE user_id=$1", notesTable)

	var notes []*GoNotes.Note

	err := r.db.Select(&notes, query, userId)

	return notes, err
}

func (r *NoteRepository) Add(userId int, note *GoNotes.Note) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description, user_id) values ($1, $2, $3) RETURNING id", notesTable)

	var noteId int

	row := r.db.QueryRow(query, note.Title, note.Description, userId)
	err := row.Scan(&noteId)

	return noteId, err
}

func (r *NoteRepository) Delete(id, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 and user_id=$2", notesTable)

	_, err := r.db.Exec(query, id, userId)

	return err
}

func (r *NoteRepository) Update(userId int, note *GoNotes.Note) (int, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if note.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, note.Title)
		argId++
	}

	if note.Description != "" {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, note.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d and user_id=$%d RETURNING id", notesTable, setQuery, argId, argId+1)
	args = append(args, note.Id, userId)

	var noteId int

	row := r.db.QueryRow(query, args...)
	err := row.Scan(&noteId)

	return noteId, err
}
