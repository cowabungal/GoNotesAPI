package repository

import (
	"GoNotes"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserInfo(id int) (*GoNotes.UserInfo, error) {
	var userInfo GoNotes.UserInfo

	query := fmt.Sprintf("SELECT username from %s WHERE id=$1", usersTable)
	err := r.db.Get(&userInfo.Username, query, id)

	return &userInfo, err
}
