package main

import (
	"log"
	"net/http"
	"os"
	"io"
	"fmt"
	"encoding/json"
	
	"github.com/FireCrackerTeam/wonderjack-web/helper"
	"github.com/FireCrackerTeam/wonderjack-web/controllers/authcontroller"
	"github.com/FireCrackerTeam/wonderjack-web/controllers/redemptioncontroller"
	"github.com/FireCrackerTeam/wonderjack-web/models"
	"github.com/gorilla/mux"
)

func appleAppSiteAssociationHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "apple-app-site-association" // Replace with the actual path to your file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/json")
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write response: %v", err), http.StatusInternalServerError)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
    name := r.URL.Query().Get("name")
    if name != "" {
        message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
    }
    fmt.Fprint(w, message)
}

func webCredHandler(w http.ResponseWriter, r *http.Request) {
    	credentials := WebCredentials{
			Apps: []string{"56YNX87MHG.com.firecracker.wonderjack"},
		}

		data := MyData{
			WebCredentials: credentials,
		}
    	helper.ResponseJSON(w, http.StatusOK, data)
}


type WebCredentials struct {
	Apps []string `json:"apps"`
}

type MyData struct {
	WebCredentials WebCredentials `json:"webcredentials"`
}

func serveWebCredentials(w http.ResponseWriter, r *http.Request) {
	fmt.Println("apple-app-site-association")
    file, err := os.Open("/home/wonderja/dev.wonderjackgame.com/.well-known/apple-app-site-association")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	  // Read the file content
	var data MyData
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Write the JSON response
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
}

func main(){
	models.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/apple-app-site-association", appleAppSiteAssociationHandler)
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/check-html", helloHandler)
	well_known := r.PathPrefix("/.well-known").Subrouter()
	well_known.HandleFunc("/apple-app-site-association", serveWebCredentials)
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/claim-redemption-code", redemptioncontroller.ClaimRedemtionCode).Methods("POST")
	// api.Use(middlewares.JWTMiddleware)
	
	// log.Fatal(http.ListenAndServe(":8080", r))
	listenAddr := ":10080"
    if val, ok := os.LookupEnv("DB_HOST"); ok {
        listenAddr = ":" + val
    }
    http.HandleFunc("/api/HttpExample", helloHandler)
    log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, r))
}
