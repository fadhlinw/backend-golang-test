package models

import "time"

type Permission struct {
	ID          int
	Slug        string
	Description string
	Active      bool
	Action      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func (t Permission) TableName() string {
	return "permission"
}
