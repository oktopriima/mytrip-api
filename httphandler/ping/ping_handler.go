package ping

import (
	"net/http"

	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/mytrip-api/context/ping"
)

type PingHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type pingHandler struct {
	rr httplib.RequestReader
	pu ping.PingUsecase
}

func NewPingHandler(rr httplib.RequestReader,
	pu ping.PingUsecase) PingHandler {
	return &pingHandler{rr, pu}
}

func (this pingHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err, resp := this.pu.Test()
	if err != nil {
		httplib.ResponseException(w, err, 422)
	}
	httplib.ResponseData(w, resp.GetData())
}
