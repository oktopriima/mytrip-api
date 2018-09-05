package repository

import "github.com/oktopriima/mytrip-api/core/model"

type UserRepository interface {
	Get(ID int64) (error, *model.Users)
	FindByEmail(Email string) (error, *model.Users)
}
