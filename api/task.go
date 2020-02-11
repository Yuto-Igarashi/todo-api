package api

type Task struct {
	ID       int       `json:id`
	Title    string    `json:title`
	Overview string    `json:overview`
	Status   string    `json:status`
	Comments []Comment `json:comments`
}
