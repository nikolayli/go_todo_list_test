package router

import (
	"todo-list/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	return r
}
