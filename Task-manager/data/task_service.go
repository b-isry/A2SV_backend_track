package data

import (
	"errors"
	"time"

	"github.com/Bisrat/task-manager/models"
)

var tasks = []models.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

type TaskService struct{}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (s *TaskService) GetAllTasks() []models.Task {
	return tasks
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func (s *TaskService) CreateTask(task models.Task) error {
	tasks = append(tasks, task)
	return nil
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) error {
	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			return nil
		}
	}
	return errors.New("task not found")
}

func (s *TaskService) DeleteTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
