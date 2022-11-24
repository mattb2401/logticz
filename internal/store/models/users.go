package models

import "time"

type User struct {
	ID       int       `gorm:"column:id"`
	Name     string    `gorm:"column:name" json:"name"`
	Username string    `gorm:"column:username" json:"username"`
	Password string    `gorm:"column:password" json:"-"`
	Added    time.Time `gorm:"column:added" json:"added"`
}

type Users []User
