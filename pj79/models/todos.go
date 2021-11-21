package models

import "time"

type Todo struct {
	ID         int
	Content    string
	UserID     int
	created_at time.Time
}
