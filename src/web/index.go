package main

import (
	"../controller"
	"github.com/gorilla/mux"
	"net/http"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/cities-suggestions", controller.SuggestionsAction).Methods("GET")
	router.HandleFunc("/cities-info", controller.InfoAction).Methods("POST")

	http.ListenAndServe(":8080", router)
}
