package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/oktopriima/mytrip-api/core/model"
	"github.com/oktopriima/mytrip-api/core/repository"
	"golang.org/x/crypto/bcrypt"
)

type GetTokenUsecase interface {
	Issue(GetTokenRequest) (error, GetTokenResponse)
	claimToken(u *model.Users) tokenResponse
}

type GetTokenRequest interface {
	GetEmail() string
	GetPassword() string
}

type GetTokenResponse interface{}

type gettokenUsecase struct {
	userRepo     repository.UserRepository
	signatureKey []byte
}

func NewGetTokenUsecase(userRepo repository.UserRepository, signatureKey []byte) GetTokenUsecase {
	return &gettokenUsecase{userRepo, signatureKey}
}

func (this gettokenUsecase) Issue(req GetTokenRequest) (error, GetTokenResponse) {
	err, user := this.CheckUser(req.GetEmail())
	if err != nil {
		return err, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))

	if err != nil {
		return errors.New("Username and password not match"), nil
	}

	at := this.claimToken(user)
	return nil, at
}

func (this gettokenUsecase) CheckUser(Email string) (error, *model.Users) {
	err, resp := this.userRepo.FindByEmail(Email)
	if err != nil {
		return err, nil
	}

	return nil, resp
}

func (this *gettokenUsecase) claimToken(u *model.Users) tokenResponse {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS512)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	expiresIn := time.Hour * (24 * 7)
	expiredAt := time.Now().Add(time.Hour * (24 * 7))
	claims["user_id"] = u.ID
	claims["exp"] = expiresIn
	refreshToken := ""

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(this.signatureKey)

	return tokenResponse{
		AccessToken:  tokenString,
		RefrashToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn.Seconds(),
		ExpiredAt:    expiredAt.Unix(),
	}
}
