package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"todo-list/models"
	"todo-list/repository"

	"github.com/gorilla/mux"
)

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body models.Task true "Task"
// @Success 201 {object} models.Task
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := task.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := repository.CreateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetTasks godoc
// @Summary Get all tasks
// @Description Get all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks [get]
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetTask godoc
// @Summary Get a task by ID
// @Description Get a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 404 {string} string "Not Found"
// @Router /tasks/{id} [get]
func GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	task, err := repository.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateTask godoc
// @Summary Update a task by ID
// @Description Update a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param task body models.Task true "Task"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	existingTask, err := checkTaskExists(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = existingTask.ID
	if err := repository.UpdateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteTask godoc
// @Summary Delete a task by ID
// @Description Delete a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	_, err = checkTaskExists(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := repository.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func checkTaskExists(id int) (models.Task, error) {
	task, err := repository.GetTask(id)
	if err != nil {
		return models.Task{}, fmt.Errorf("task not found")
	}
	return task, nil
}
