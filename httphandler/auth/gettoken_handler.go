package auth

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/mytrip-api/context/auth"
)

type GetTokenHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type gettokenHandler struct {
	tuc auth.GetTokenUsecase
	rr  httplib.RequestReader
}

func NewGetTokenHandler(tuc auth.GetTokenUsecase, rr httplib.RequestReader) GetTokenHandler {
	return &gettokenHandler{tuc, rr}
}

func (this gettokenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := gettokenRequest{}
	err := this.rr.GetJsonData(r, &req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	err, resp := this.tuc.Issue(req)
	if err != nil {
		httplib.ResponseException(w, err, 422)
		return
	}
	httplib.ResponseData(w, resp)
}
