package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      *string    `json:"name"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	Birthdate *time.Time `json:"birthdate"`
	Gender    string     `gorm:"default:'other'" json:"gender"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
