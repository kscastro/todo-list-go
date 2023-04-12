package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kscastro/todo-list-go/src/database"
	"github.com/kscastro/todo-list-go/src/model"
)

type TaskController struct {
	db database.TaskRepository
}

func NewTaskController(db database.TaskRepository) (*TaskController, error) {
	if db == nil {
		return nil, fmt.Errorf("Nil Repository")
	}

	return &TaskController{db}, nil
}

func (t TaskController) GetAllTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	payload := t.db.GetAllTask()
	json.NewEncoder(w).Encode(payload)
}

func (t TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var task model.TodoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	t.db.InsertOneTask(task)
	json.NewEncoder(w).Encode(task)
}

func (t TaskController) TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	t.db.TaskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func (t TaskController) UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	t.db.UndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func (t TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	t.db.DeleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func (t TaskController) DeleteAllTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	count := t.db.DeleteAllTask()
	json.NewEncoder(w).Encode(count)

}
