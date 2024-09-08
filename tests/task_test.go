package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-list/config"
	"todo-list/models"
	"todo-list/repository"
	"todo-list/router"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	config.LoadConfig()
	r := router.SetupRouter()

	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTasks(t *testing.T) {
	config.LoadConfig()
	r := router.SetupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTask(t *testing.T) {
	config.LoadConfig()
	r := router.SetupRouter()

	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	if err := repository.CreateTask(&task); err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	config.LoadConfig()
	r := router.SetupRouter()

	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	if err := repository.CreateTask(&task); err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	updatedTask := models.Task{
		Title:       "Updated Task",
		Description: "This is an updated test task",
	}

	jsonValue, _ := json.Marshal(updatedTask)
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	config.LoadConfig()
	r := router.SetupRouter()

	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}
	if err := repository.CreateTask(&task); err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
