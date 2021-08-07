package models

import (
	"time"
)

type Note struct {
	ID uint64 `gorm:"primary_key;auto_increment;" json:"id"`
	Title string `gorm:"not null;" json:"title"`
	Content string `json:"content"`
	CategoryId uint `json:"category_id"`
	Category Category `gorm:"foreignKey:CategoryId" json:"category"`
	UserID uint `json:"user_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
