package controllers

import (
	"net/http"
	"strings"

	"github.com/Bisrat/task-manager/data"
	"github.com/Bisrat/task-manager/models"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *data.TaskService
}

func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks := c.taskService.GetAllTasks()
	ctx.JSON(http.StatusOK, models.APIResponse{
		Message: "Tasks retrieved successfully",
		Data:    tasks,
	})
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Message: "Task ID is required",
		})
		return
	}

	task, err := c.taskService.GetTaskByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.APIResponse{
			Message: "Task not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Message: "Task retrieved successfully",
		Data:    task,
	})
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Message: "Invalid request payload",
			Data: []models.ValidationError{
				{Field: "body", Message: err.Error()},
			},
		})
		return
	}

	var validationErrors []models.ValidationError
	if newTask.Title == "" {
		validationErrors = append(validationErrors, models.ValidationError{
			Field:   "title",
			Message: "Title is required",
		})
	}
	if newTask.Status != "" && !isValidStatus(newTask.Status) {
		validationErrors = append(validationErrors, models.ValidationError{
			Field:   "status",
			Message: "Status must be one of: pending, in_progress, completed",
		})
	}

	if len(validationErrors) > 0 {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Message: "Validation failed",
			Data:    validationErrors,
		})
		return
	}

	if newTask.Status == "" {
		newTask.Status = models.StatusPending
	}

	if err := c.taskService.CreateTask(newTask); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.APIResponse{
			Message: "Failed to create task",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.APIResponse{
		Message: "Task created successfully",
		Data:    newTask,
	})
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Message: "Task ID is required",
		})
		return
	}

	var updatedTask models.Task
	if err := ctx.BindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Message: "Invalid request payload",
			Data: []models.ValidationError{
				{Field: "body", Message: err.Error()},
			},
		})
		return
	}

	if updatedTask.Status != "" && !isValidStatus(updatedTask.Status) {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Message: "Validation failed",
			Data: []models.ValidationError{
				{
					Field:   "status",
					Message: "Status must be one of: pending, in_progress, completed",
				},
			},
		})
		return
	}

	if err := c.taskService.UpdateTask(id, updatedTask); err != nil {
		ctx.JSON(http.StatusNotFound, models.APIResponse{
			Message: "Task not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Message: "Task updated successfully",
	})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.APIResponse{
			Message: "Task ID is required",
		})
		return
	}

	if err := c.taskService.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, models.APIResponse{
			Message: "Task not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.APIResponse{
		Message: "Task deleted successfully",
	})
}

func isValidStatus(status string) bool {
	status = strings.ToLower(status)
	return status == models.StatusPending ||
		status == models.StatusInProgress ||
		status == models.StatusCompleted
}
