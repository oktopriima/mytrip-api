package service

import (
	"database/sql"
	"encoding/json"

	"github.com/oktopriima/mytrip-api/core/model"
	"github.com/oktopriima/mytrip-api/core/repository"
)

type userService struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userService{db}
}

func (this userService) Get(ID int64) (error, *model.Users) {
	m := new(model.Users)
	query := "SELECT id, name, email, remember_token, active, activation_token, created_at, updated_at, deleted_at FROM users WHERE deleted_at IS NULL AND id = ?"

	row := this.db.QueryRow(query, ID)
	err := row.Scan(&m.ID, &m.Name, &m.Email, &m.RememberToken, &m.Active, &m.ActiveToken, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)

	if err != nil {
		return err, nil
	}
	_, err = json.Marshal(&m)

	if err != nil {
		return err, nil
	}
	return nil, m
}

func (this userService) FindByEmail(Email string) (error, *model.Users) {
	m := new(model.Users)
	query := "SELECT id, name, email, password" +
		"	 FROM users" +
		"	WHERE email = ? "
	row := this.db.QueryRow(query, Email)
	err := row.Scan(&m.ID, &m.Name, &m.Email, &m.Password)
	if err != nil {
		return err, nil
	}

	return nil, m
}
