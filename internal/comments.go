package internal

import (
	"log"
	"todo-api/api"
	"todo-api/config"

	"github.com/gin-gonic/gin"
)

//TaskIdに一致したコメントを取得する
func GetComment(TaskId int) []api.Comment {
	db := config.GormConnect()
	defer db.Close()
	commnets := []api.Comment{}
	db.Where("task_id = ?", TaskId).Find(&commnets)
	return commnets
}

func AddComment(ctx *gin.Context) {
	//id := ctx.Param("task_id")
	db := config.GormConnect()
	defer db.Close()
	log.Println("addComment start")
	comment := api.Comment{}
	//comment.CeateDate = time.Now()
	if err := ctx.BindJSON(&comment); err != nil {
		log.Println(err)
		return
	}

	//comment.TaskId, _ = strconv.Atoi(id)
	db.Create(&comment)
	log.Println(db)
	ctx.JSON(200, comment)
}
