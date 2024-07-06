package models

import "time"

type Task struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	TotalHours  float64   `json:"total_hours"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Workload struct {
	TaskName   string  `json:"taskName"`
	TotalHours float64 `json:"totalHours"`
}
