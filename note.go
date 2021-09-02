package GoNotes

import "errors"

type Note struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UserId      string `json:"user_id" db:"user_id"`
}

func (i Note) Validate() error {
	if i.Title == "" && i.Description == "" {
		return errors.New("note structure has no values")
	}

	return nil
}
