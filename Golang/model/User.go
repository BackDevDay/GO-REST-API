package model

import "time"

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"unique" form:"email" binding:"required"`
	Password  string `form:"password" binding:"required"`
	CreatedAt time.Time
}
