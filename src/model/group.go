package model

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupID     string
	Name        string
	Description string
	OwnerID     string
}

type Member struct {
	gorm.Model
	GroupID  string
	MemberID string
}

type GroupTask struct {
	gorm.Model
	GroupTaskID string
	Name        string
	Type        string
	Content     string
	DeadLine    string
	IsComplate  bool
}
