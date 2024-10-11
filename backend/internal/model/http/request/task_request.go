package request

import "time"

type TaskCreateRequest struct {
	Content    string    `json:"content" binding:"required"`
	Deadline   time.Time `json:"deadline"`
	AssigneeID string    `json:"assignee_id"`
}
