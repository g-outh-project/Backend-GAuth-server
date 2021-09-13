package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string
	Password string
	Name     string
	Nickname string
	Birth    string
	Email    string
	School   string
}
