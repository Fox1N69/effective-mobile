package models

import "time"

type Task struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	Name       string    `json:"name"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	TotalHours float64   `json:"total_hours"`
}
