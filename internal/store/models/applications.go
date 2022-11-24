package models

import "time"

type Application struct {
	ID    int       `gorm:"column:id"`
	Name  string    `gorm:"column:name"`
	Added time.Time `gorm:"column:added"`
}

type Applications []Application
