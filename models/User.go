package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint      `gorm:"primary_key;auto_increment" json:"id"`
	Username   string    `gorm:"size:255;not null;unique" json:"username"`
	Email      string    `gorm:"size:100;not null;unique" json:"email"`
	Password   string    `gorm:"size:100;not null;" json:"password"`
	AvatarPath string    `gorm:"size:255;null;" json:"avatar_path"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) CreateUser(db *gorm.DB) (*User, error) {

	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
