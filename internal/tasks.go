package internal

import (
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
	ctx.JSON(200, tasks)
}

func GetTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	db := config.GormConnect()
	var task api.Task
	idInt, _ := strconv.Atoi(id)
	task.Comments = GetComment(idInt)
	err := db.First(&task, id).Error
	notExistId(ctx, err, id)
	if err == nil {
		ctx.JSON(200, task)
	}
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
		log.Println("Create Tasks :Missing input")
		return
	}
	db.Create(&task)
	log.Println(db)
	ctx.JSON(200, task)
}

func UpdateTasksById(ctx *gin.Context) {
	id := ctx.Param("id")
	db := config.GormConnect()
	task := api.Task{}
	task.ID, _ = strconv.Atoi(id)
	updateTask := task
	if err := ctx.BindJSON(&updateTask); err != nil {
		log.Println(err)
		return
	}
	if updateTask.Title == "" || updateTask.Overview == "" || updateTask.Status == "" {
		log.Println("Update Tasks :Missing input")
		return
	}
	err := db.First(&task, id).Error
	notExistId(ctx, err, id)
	if err != nil {
		return
	}
	db.Save(&updateTask)
	ctx.JSON(200, updateTask)
}

func DeleteTasksById(ctx *gin.Context) {
	log.Println("DeleteTasksById start")
	id := ctx.Param("id")
	db := config.GormConnect()
	task := api.Task{}
	task.ID, _ = strconv.Atoi(id)
	defer db.Close()
	db.First(&task, id)
	db.Delete(&task)
}

//id がない時のエラー関数
func notExistId(ctx *gin.Context, err interface{}, id string) {
	if err != nil {
		log.Printf("ID:[%s] does not exist", id)
		log.Println(err)
		ctx.JSON(200, gin.H{
			"staus": "NotFound",
		})
	}
}
