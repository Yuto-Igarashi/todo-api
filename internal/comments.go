package internal

import (
	"todo-api/api"
	"todo-api/config"
)

//TaskIdに一致したコメントを取得する
func GetComment(TaskId int) []api.Comment {
	db := config.GormConnect()
	defer db.Close()
	commnets := []api.Comment{}
	db.Where("task_id = ?", TaskId).Find(&commnets)
	return commnets
}
