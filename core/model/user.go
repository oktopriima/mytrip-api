package model

import (
	"time"

	"github.com/oktopriima/go-package/library"
)

type Users struct {
	ID            int64              `json:"id"`
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	Password      string             `json:"password"`
	RememberToken library.NullString `json:"remember_token"`
	Active        int                `json:"active"`
	ActiveToken   library.NullString `json:"activation_token"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	DeletedAt     library.NullTime   `json:"deleted_at"`
}
