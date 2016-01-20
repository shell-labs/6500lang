package model

import "github.com/jinzhu/gorm"

type Bible struct {
	gorm.Model
	BibleID     string
	Name        string
	Author      string
	Description string
	WrittenAt   string
	License     string
	IsOrignal   bool
}

type Book struct {
	gorm.Model
	BookID       string
	Name         string
	ShortName    string
	BookNumber   uint
	Author       string
	Description  string
	ChapterVerse string // object{int: array} {1:[1,2,3...], 2:[1,2,3...]}
}

type Verse struct {
	gorm.Model
	VerseID       string
	BibleID       string
	BookNumber    uint
	ChapterNumber uint
	VerseNumber   uint
	Text          string
	FootNote      string
}

type Word struct {
	gorm.Model
	VerseID      string
	Word         string
	StrongNumber string
	Comment      string
}

type Vocabulary struct {
	gorm.Model
	StrongNumber string
	Mean         string
	Protype      string
	Morphology   string
	Defition     string
	Comment      string
}
