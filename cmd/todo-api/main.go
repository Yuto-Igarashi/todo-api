package main

import (
	"todo-api/internal/tasks"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/tasks", tasks.GetTasks)
		api.POST("/tasks", tasks.CreateTask)
		api.GET("/tasks/:id", tasks.GetTasksById)
		api.POST("/tasks/:id", tasks.UpdateTasksById)
	}
	router.Run(":8080")
}
