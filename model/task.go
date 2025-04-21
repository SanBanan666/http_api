package model

type TaskStatus string

const (
	StatusPending TaskStatus = "pending"
	StatusRunning TaskStatus = "running"
	StatusSuccess TaskStatus = "success"
	StatusFailed  TaskStatus = "failed"
)

type Task struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
	Result *string    `json:"result"`
	Error  *string    `json:"error"`
}
