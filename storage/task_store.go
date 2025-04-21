package storage

import (
	"http_api/model"
	"sync"
)

var (
	taskStore = make(map[string]*model.Task)
	mu        sync.RWMutex
)

func CreateTask(id string) {
	mu.Lock()
	defer mu.Unlock()
	taskStore[id] = &model.Task{
		ID:     id,
		Status: model.StatusPending,
	}
}

func UpdateTaskSatus(id string, status model.TaskStatus, result *string, err *string) {
	mu.Lock()
	defer mu.Unlock()
	if task, exitts := taskStore[id]; exitts {
		task.Status = status
		task.Result = result
		task.Error = err
	}
}

func GetTask(id string) (*model.Task, bool) {
	mu.RLock()
	defer mu.RUnlock()
	task, exists := taskStore[id]
	return task, exists
}
