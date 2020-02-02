package tasks

import (
	"fmt"
	"log"
	"strconv"
	"todo-api/api"
	"todo-api/config"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	db := config.GormConnect()
	defer db.Close()
	tasks := []api.Task{}
	db.Find(&tasks)
	fmt.Println(tasks)
	ctx.JSON(200, tasks)
}

func GetTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	db := config.GormConnect()
	var task api.Task
	db.First(&task, id)
	ctx.JSON(200, task)
}

func CreateTask(ctx *gin.Context) {
	db := config.GormConnect()
	log.Println("CreateTask start")
	task := api.Task{}
	if err := ctx.BindJSON(&task); err != nil {
		log.Println(err)
		return
	}
	if task.Title == "" || task.Overview == "" || task.Status == "" {
		return
		//エラーを返すようにしたい
	}
	db.Create(&task)
	fmt.Println(db)
	ctx.JSON(200, task)
}

func UpdateTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	db := config.GormConnect()
	task := api.Task{}
	task.ID, _ = strconv.Atoi(id)
	updateTask := task
	if err := ctx.BindJSON(&updateTask); err != nil {
		return
	}
	if updateTask.Title == "" || updateTask.Overview == "" || updateTask.Status == "" {
		return
	}
	db.Save(&updateTask)
	ctx.JSON(200, updateTask)
}
