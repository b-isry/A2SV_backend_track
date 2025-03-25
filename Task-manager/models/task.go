package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}


const (
	StatusPending    = "pending"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
)
