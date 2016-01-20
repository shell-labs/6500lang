package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	TaskID     string
	Name       string
	OwnerID    string
	Type       string
	Content    string
	DeadLine   string
	IsComplate bool
}
