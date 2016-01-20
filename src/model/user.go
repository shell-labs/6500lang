package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string
	UserName string
	Password string
	Email    string
	Mobile   string
	Provide  string
}

type Profile struct {
	gorm.Model
	UserID   string
	NickName string
	Avatar   string
	HomeURL  string
	Location string
}
