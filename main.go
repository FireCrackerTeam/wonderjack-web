package main

import (
	"log"
	"net/http"

	"github.com/FireCrackerTeam/wonderjack-web/controllers/authcontroller"
	"github.com/FireCrackerTeam/wonderjack-web/controllers/testcontroller"
	"github.com/FireCrackerTeam/wonderjack-web/middlewares"
	"github.com/FireCrackerTeam/wonderjack-web/models"
	"github.com/gorilla/mux"
)

func main(){
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/test", testcontroller.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}