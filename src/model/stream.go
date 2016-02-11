package model

import "github.com/jinzhu/gorm"

// Stream 信息流
// TODO 同步到ES
type Stream struct {
	gorm.Model
	UserID  string
	Content string
	Type    string
	Cover   string
	URL     string
}
