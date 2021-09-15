package model

import "time"

type User struct {
	Id        string `gorm:"primaryKey"`
	Password  string
	Name      string
	Nickname  string
	Birth     string
	Email     string
	School    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
