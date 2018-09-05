package auth

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/mytrip-api/context/auth"
	"github.com/oktopriima/mytrip-api/core/service"
)

func NewRoute(api *mux.Router, rr httplib.RequestReader, db *sql.DB, signatureKey []byte) {
	userRepo := service.NewUserRepository(db)

	tuc := auth.NewGetTokenUsecase(userRepo, signatureKey)
	suc := auth.NewSelfUsecase(userRepo)
	guauc := auth.NewGetAuthUserUsecase()

	gth := NewGetTokenHandler(tuc, rr)
	sh := NewSelfHandler(suc, rr, guauc)

	authRoute := api.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("", gth.Handle).Methods("POST")
	authRoute.HandleFunc("", httplib.JWTMid(sh.Handle)).Methods("GET")
}
