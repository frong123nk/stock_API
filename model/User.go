package model

import "time"

type User struct {
	ID        int    `gorm:"primary_key"`
	Username  string `gorm:"unique" form:"username" binding:"required"`
	Password  string `from:"password" binding:"required"`
	Level     string `gorm:"default:normal"`
	CreatedAt time.Time
}
