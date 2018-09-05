package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mmuflih/go-httplib/httplib"
	"github.com/oktopriima/mytrip-api/conf"
	"github.com/oktopriima/mytrip-api/httphandler/auth"
	"github.com/oktopriima/mytrip-api/httphandler/ping"
	"github.com/oktopriima/mytrip-api/lib"
)

var signingKey = []byte("JWTtoken")
var config conf.Config

func init() {
	config = conf.NewConfig()
	fmt.Println()
	fmt.Println("   _____    __  ___  ________    _____")
	fmt.Println("  / ___ \\  |  |/  / |__    __|  / ___ \\ ")
	fmt.Println(" / /   \\ \\ |     /     |  |    / /   \\ \\ ")
	fmt.Println(" \\ \\___/ / |     \\     |  |    \\ \\___/ / ")
	fmt.Println("  \\_____/  |__|\\__\\    |__|     \\_____/ ")
}

func main() {
	mydb := mysqlConn()

	muxrr := httplib.NewMuxRequestReader()
	httplib.InitJWTMiddleware(signingKey)

	api := mux.NewRouter()
	apivi := api.PathPrefix("/api/v1").Subrouter()

	/** list route */
	ping.NewRoute(apivi, muxrr, mydb)
	auth.NewRoute(apivi, muxrr, mydb, signingKey)

	/** cors */
	headersVal := config.GetStrings("cors.allowed_headers")
	methodsVal := config.GetStrings("cors.allowed_methods")
	originsVal := config.GetStrings("cors.allowed_origins")
	headers := handlers.AllowedHeaders(headersVal)
	methods := handlers.AllowedMethods(methodsVal)
	origins := handlers.AllowedOrigins(originsVal)
	cors := handlers.CORS(headers, methods, origins)

	log.Println("------------------------------------------------------")
	log.Println("Last Build at "+time.Now().Format("2006-01-02 15:04:05"), "Listen on port", config.GetString("server.address"))
	log.Fatal(http.ListenAndServe(config.GetString(`server.address`), cors(lib.Logger(api))))
}

func mysqlConn() *sql.DB {
	err, db := conf.GetMysqlConn(config)
	if err != nil {
		fmt.Println("Mysql error", err.Error())
	}
	return db
}
