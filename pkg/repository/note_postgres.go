package repository

import (
	"GoNotes"
	"fmt"
	"github.com/jmoiron/sqlx"
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

func (r *NoteRepository) Add(id int, note *GoNotes.Note) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, description, user_id) values ($1, $2, $3) RETURNING id", notesTable)

	var noteId int

	row := r.db.QueryRow(query, note.Title, note.Description, id)
	err := row.Scan(&noteId)

	return noteId, err
}

func (r *NoteRepository) Delete(id, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 and user_id=$2", notesTable)

	_, err := r.db.Exec(query, id, userId)

	return err
}
