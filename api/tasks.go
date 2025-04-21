package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"http_api/service"
	"http_api/storage"
	"net/http"
)

func RegisterRoutes(r chi.Router) {
	r.Post("/tasks", createTaskHandler)
	r.Get("/tasks/{taskID}", getTaskHandler)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := service.CreateTask()

	resp := map[string]string{
		"task_id": taskID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	task, found := storage.GetTask(taskID)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
