package api

type Comment struct {
	ID        int    `json:id`
	Content   string `json:content`
	CeateDate string `json:ceate_date`
	TaskId    int    `json:task_id`
}
