package auth

import "github.com/oktopriima/mytrip-api/core/model"

type selfResponse struct {
	Data *model.Users
}

func (this selfResponse) GetResponse() *model.Users {
	return this.Data
}
