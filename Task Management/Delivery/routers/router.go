package routers

import (
	"Task-Management/Delivery/controllers"
	infrastructure "Task-Management/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController, taskController *controllers.TaskController) *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/register", userController.Register)
		public.POST("/login", userController.Login)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(infrastructure.AuthMiddleware())
	{
		// User routes
		protected.GET("/users", userController.GetAllUsers)

		// Task routes
		protected.POST("/tasks", taskController.CreateTask)
		protected.GET("/tasks", taskController.GetTasksByUserID)
		protected.GET("/tasks/:id", taskController.GetTaskByID)
		protected.PUT("/tasks/:id", taskController.UpdateTask)
		protected.DELETE("/tasks/:id", taskController.DeleteTask)
	}

	// Admin routes
	admin := router.Group("/api/admin")
	admin.Use(infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware())
	{
		admin.GET("/tasks", taskController.GetAllTasks)
	}

	return router
}
