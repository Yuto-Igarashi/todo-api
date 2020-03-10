package main

import (
	"todo-api/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/tasks", internal.GetTasks)
		api.POST("/tasks", internal.CreateTask)
		api.GET("/tasks/:id", internal.GetTasksById)
		api.POST("/tasks/:id", internal.UpdateTasksById)
		api.DELETE("/tasks/:id", internal.DeleteTasksById)

		//コメント
		api.POST("/comment/", internal.AddComment)
	}
	router.NoRoute(internal.NotFoundRoute)
	router.Run(":8080")
}
