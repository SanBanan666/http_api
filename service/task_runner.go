package service

import (
	"github.com/google/uuid"
	"http_api/model"
	"http_api/storage"
	"time"
)

var taskQueue = make(chan string, 100)

func StartWorker() {
	go func() {
		for taskID := range taskQueue {
			storage.UpdateTaskSatus(taskID, model.StatusRunning, nil, nil)

			// Имитируем длительную обработку
			time.Sleep(3 * time.Minute)

			result := "Result for task " + taskID
			storage.UpdateTaskSatus(taskID, model.StatusSuccess, &result, nil)
		}
	}()
}

func CreateTask() string {
	id := uuid.New().String()
	storage.CreateTask(id)
	taskQueue <- id
	return id
}
