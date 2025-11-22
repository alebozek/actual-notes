package models

import "time"

type Session struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
