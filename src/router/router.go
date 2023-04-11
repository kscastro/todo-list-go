package router

import (
	"github.com/gorilla/mux"
	"github.com/kscastro/todo-list-go/src/controller"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/", controller.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/task/complete/{id}", controller.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/undo/{id}", controller.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/delete{id}", controller.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/task/deleteAll", controller.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
