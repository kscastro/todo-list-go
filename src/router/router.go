package router

import (
	"github.com/gorilla/mux"
	"github.com/kscastro/todo-list-go/src/database"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", database.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/", database.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/task/complete/{id}", database.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/undo/{id}", database.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/delete{id}", database.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/task/deleteAll", database.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
