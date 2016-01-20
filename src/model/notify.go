package model

import "github.com/jinzhu/gorm"

type Notify struct {
	gorm.Model
	NotifyID string
	FromID   string
	ToID     string
	URL      string
	Content  string
	IsReaded bool
	ReadAt   string
}
