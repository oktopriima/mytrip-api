package auth

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/mytrip-api/context/auth"
)

type SelfHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type selfHandler struct {
	suc  auth.SelfUsecase
	rr   httplib.RequestReader
	auth auth.GetAuthUserUsecase
}

func NewSelfHandler(suc auth.SelfUsecase, rr httplib.RequestReader, auth auth.GetAuthUserUsecase) SelfHandler {
	return &selfHandler{suc, rr, auth}
}

func (this selfHandler) Handle(w http.ResponseWriter, r *http.Request) {
	UserID := this.auth.GetUserID(r)
	err, resp := this.suc.Index(UserID)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}

	httplib.ResponseData(w, resp.GetResponse())
}
