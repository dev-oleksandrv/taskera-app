package response

import "github.com/google/uuid"

type TaskDto struct {
	ID          uuid.UUID `json:"id"`
	Content     string    `json:"content"`
	FullContent string    `json:"full_content"`
	Completed   bool      `json:"completed"`
	Order       int       `json:"order"`
}

type TaskGetAllResponse struct {
	Tasks []TaskDto `json:"tasks"`
}

type TaskCreateResponse struct {
	Task TaskDto `json:"task"`
}
