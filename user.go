package GoNotes

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
}
