package model

import (
	"time"

	"github.com/oktopriima/mytrip-api/lib"
)

type Users struct {
	ID            int64          `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	RememberToken lib.NullString `json:"remember_token"`
	Active        int            `json:"active"`
	ActiveToken   lib.NullString `json:"activation_token"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     lib.NullTime   `json:"deleted_at"`
}
