package model

import "github.com/jinzhu/gorm"

type Stream struct {
	gorm.Model
	UserID  string
	Content string
	Type    string
	Cover   string
	URL     string
}
