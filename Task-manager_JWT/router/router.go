package router

import (
	"github.com/Bisrat/task-manager/controllers"
	"github.com/Bisrat/task-manager/data"
	"github.com/Bisrat/task-manager/middleware"
	"github.com/Bisrat/task-manager/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userController := controllers.NewUserController()
	taskService := data.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/users", middleware.RequireRole(models.RoleAdmin), userController.GetAllUsers)
		authorized.POST("/tasks", taskController.CreateTask)
		authorized.GET("/tasks", taskController.GetTasks)
		authorized.GET("/tasks/:id", taskController.GetTaskByID)
		authorized.PUT("/tasks/:id", taskController.UpdateTask)
		authorized.DELETE("/tasks/:id", taskController.DeleteTask)
	}

	return router
}
