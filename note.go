package GoNotes

type Note struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UserId      string `json:"user_id" db:"user_id"`
}
