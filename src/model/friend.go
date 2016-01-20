package model

import "github.com/jinzhu/gorm"

type Friend struct {
	gorm.Model
	FriendID   string
	FollowerID string
	FriendName string
}
