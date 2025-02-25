package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

/*
type Book struct 			- এটি একটি Go Struct, যা MySQL-এ টেবিল তৈরি করবে।
gorm.Model 					- এটি ID, CreatedAt, UpdatedAt, DeletedAt ফিল্ড যোগ করবে।
Title string json:"title" 	- JSON API Response যেন title নামে আসে।
*/
