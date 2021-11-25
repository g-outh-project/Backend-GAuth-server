package model

import "time"

type Client struct {
	ClientId  string `gorm:"primaryKey"`
	Secret    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
