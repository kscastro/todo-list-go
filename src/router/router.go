package router

import (
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", "metodos").Methods("GET", "OPTIONS")
	router.HandleFunc("/", "metodos").Methods("POST", "OPTIONS")
	router.HandleFunc("/task/complete/{id}", "metodos").Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/undo/{id}", "metodos").Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/delete{id}", "metodos").Methods("DELETE", "OPTIONS")
	router.HandleFunc("/task/deleteAll", "metodos").Methods("DELETE", "OPTIONS")
	return router
}
