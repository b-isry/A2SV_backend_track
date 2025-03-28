package controllers

import (
	"net/http"

	domain "Task-Management/Domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	userUseCase domain.UserUseCase
}

type TaskController struct {
	taskUseCase domain.TaskUseCase
}

func NewUserController(userUseCase domain.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func NewTaskController(taskUseCase domain.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}

// User Controllers
func (c *UserController) Register(ctx *gin.Context) {
	var req domain.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	createdUser, err := c.userUseCase.Register(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Message: "User registered successfully",
		Data:    createdUser,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var req domain.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	user, token, err := c.userUseCase.Login(ctx.Request.Context(), req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Message: "Login successful",
		Data: gin.H{
			"token": token,
			"user":  user,
		},
	})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userUseCase.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Message: "Users retrieved successfully",
		Data:    users,
	})
}

// Task Controllers
func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	// Get user ID from context
	userID, _ := ctx.Get("user_id")
	id, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: "Invalid user ID"})
		return
	}
	task.UserID = id

	createdTask, err := c.taskUseCase.CreateTask(ctx.Request.Context(), &task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, domain.APIResponse{
		Message: "Task created successfully",
		Data:    createdTask,
	})
}

func (c *TaskController) GetTaskByID(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: "Invalid task ID"})
		return
	}

	task, err := c.taskUseCase.GetTaskByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Message: "Task retrieved successfully",
		Data:    task,
	})
}

func (c *TaskController) GetTasksByUserID(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	id, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: "Invalid user ID"})
		return
	}

	tasks, err := c.taskUseCase.GetTasksByUserID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Message: "Tasks retrieved successfully",
		Data:    tasks,
	})
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.taskUseCase.GetAllTasks(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Message: "Tasks retrieved successfully",
		Data:    tasks,
	})
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: "Invalid task ID"})
		return
	}

	task.ID = id
	if err := c.taskUseCase.UpdateTask(ctx.Request.Context(), &task); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Message: "Task updated successfully",
	})
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: "Invalid task ID"})
		return
	}

	if err := c.taskUseCase.DeleteTask(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Message: "Task deleted successfully",
	})
}
