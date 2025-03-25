package router

import (
	"github.com/Bisrat/task-manager/controllers"
	"github.com/Bisrat/task-manager/data"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	taskService := data.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	tasks := router.Group("/tasks")
	{
		tasks.GET("", taskController.GetTasks)
		tasks.GET("/:id", taskController.GetTaskByID)
		tasks.POST("", taskController.CreateTask)
		tasks.PUT("/:id", taskController.UpdateTask)
		tasks.DELETE("/:id", taskController.DeleteTask)
	}

	return router
}
