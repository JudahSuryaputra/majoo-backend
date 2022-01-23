package http

import (
	"log"
	"majoo-backend/http/Auth"
	"majoo-backend/http/User"
	"majoo-backend/http/middlewares"
	"net/http"
	"os"

	"github.com/gocraft/dbr"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(dbConn *dbr.Connection) {
	a.Router = mux.NewRouter().StrictSlash(true)

	a.authRoutes(dbConn)
	a.userRoutes(dbConn)
}

func (a *App) authRoutes(dbConn *dbr.Connection) {
	a.Router.Use(middlewares.SetContentTypeMiddleware)

	login := Auth.Login{DBConn: dbConn}
	a.Router.Handle("/v1/login", login).Methods(http.MethodPost)
}

func (a *App) userRoutes(dbConn *dbr.Connection) {
	user := a.Router.PathPrefix("/v1/api").Subrouter()
	user.Use(middlewares.CommonAuthJwtVerify)

	getMerchantTransactions := User.GetMerchantTransactions{DBConn: dbConn}
	getOutletTransactions := User.GetOutletTransactions{DBConn: dbConn}
	user.Handle("/merchant/transactions", getMerchantTransactions).Methods(http.MethodGet)
	user.Handle("/outlet/transactions/{id}", getOutletTransactions).Methods(http.MethodGet)
}

func (a *App) RunServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("PORT")
	}
	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"})

	log.Printf("\nServer starting on Port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CombinedLoggingHandler(os.Stderr, handlers.CORS(headersOK, originsOK, methodsOK)(a.Router))))
}
