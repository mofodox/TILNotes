package models

import (
	"time"
)

type Note struct {
	ID uint64 `gorm:"primary_key;auto_increment;" json:"id"`
	Title string `gorm:"not null;" json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
