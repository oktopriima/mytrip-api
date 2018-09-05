package auth

import (
	"github.com/oktopriima/mytrip-api/core/model"
	"github.com/oktopriima/mytrip-api/core/repository"
)

type SelfUsecase interface {
	Index(UserID int64) (error, SelfResponse)
}

type SelfResponse interface {
	GetResponse() *model.Users
}

type selfUsecase struct {
	uRepo repository.UserRepository
}

func NewSelfUsecase(uRepo repository.UserRepository) SelfUsecase {
	return &selfUsecase{uRepo}
}

func (this selfUsecase) Index(UserID int64) (error, SelfResponse) {
	err, resp := this.uRepo.Get(UserID)
	if err != nil {
		return err, selfResponse{}
	}

	return nil, selfResponse{resp}
}
