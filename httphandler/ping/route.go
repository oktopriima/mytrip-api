package ping

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/mytrip-api/context/ping"
)

func NewRoute(api *mux.Router, rr httplib.RequestReader, db *sql.DB) {
	/** Usecase */
	pu := ping.NewPingUsecase()

	/** Handler */
	ph := NewPingHandler(rr, pu)

	pingRoute := api.PathPrefix("/ping").Subrouter()
	pingRoute.HandleFunc("", ph.Handle).Methods("GET")
}
