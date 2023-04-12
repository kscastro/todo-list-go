package router

import (
	"github.com/gorilla/mux"
	"github.com/kscastro/todo-list-go/src/controller"
	"github.com/kscastro/todo-list-go/src/database"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	db := database.NewDB()

	ctr, err := controller.NewTaskController(db)
	if err != nil {
		return nil
	}

	router.HandleFunc("/", ctr.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/", ctr.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/task/complete/{id}", ctr.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/undo/{id}", ctr.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/task/delete/{id}", ctr.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/task/deleteAll", ctr.DeleteAllTask).Methods("DELETE", "OPTIONS")
	return router
}
