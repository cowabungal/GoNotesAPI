package repository

import (
	"GoNotes"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user GoNotes.User) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", usersTable)

	// набор аргументов подставляется в плейсхолдеры в виде доллара, возвращает объект роу, который хранит инфу о возвращаемой строке из базы
	// В нашем случае возвращаем одну строку с полем значения id
	row := r.db.QueryRow(query, user.Username, user.Password)
	err := row.Scan(&user.Id)

	return err
}

func (r *AuthRepository) CreateSession(user *GoNotes.User) (*GoNotes.User, error) {
	user, err := r.CheckUser(user.Username, user.Password)

	return user, err
}

func (r *AuthRepository) CheckUser(username, password string) (*GoNotes.User, error) {
	var user GoNotes.User

	query := fmt.Sprintf("SELECT id from %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return &user, err
}

func (r *AuthRepository) CheckSession(id int) error {
	var username string

	query := fmt.Sprintf("SELECT username from %s WHERE id=$1", usersTable)
	err := r.db.Get(&username, query, id)

	return err
}
