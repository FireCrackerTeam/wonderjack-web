package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/FireCrackerTeam/wonderjack-web/controllers/authcontroller"
	"github.com/FireCrackerTeam/wonderjack-web/controllers/redemptioncontroller"
	"github.com/FireCrackerTeam/wonderjack-web/middlewares"
	"github.com/FireCrackerTeam/wonderjack-web/models"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
    name := r.URL.Query().Get("name")
    if name != "" {
        message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
    }
    fmt.Fprint(w, message)
}

func main(){
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/claim-redemption-code", redemptioncontroller.ClaimRedemtionCode).Methods("POST")
	api.Use(middlewares.JWTMiddleware)

	// log.Fatal(http.ListenAndServe(":80", r))
	listenAddr := ":80"
    if val, ok := os.LookupEnv("DB_HOST"); ok {
        listenAddr = ":" + val
    }
    http.HandleFunc("/api/HttpExample", helloHandler)
    log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, r))
}
