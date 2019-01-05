package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/cikupin/testing_code/filtering_input/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.GoOzzoValidator).Methods("POST")

	serve := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Println("Service is running on port 8080..")
	serve.ListenAndServe()
}
