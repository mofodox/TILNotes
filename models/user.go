package models

import "time"

type User struct {
	ID uint `gorm:"primary_key;auto_increment;" json:"id"`
	Email string `gorm:"not null;unique;" json:"email"`
	FirstName string `gorm:"not null;" json:"first_name"`
	LastName string `gorm:"not null;" json:"last_name"`
	Password []byte `gorm:"not null;" json:"-"`
	Notes []Note `json:"notes"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
