package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title string `gorm:"not null;"`
	Content string
}
