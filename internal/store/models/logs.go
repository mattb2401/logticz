package models

import "time"

type Log struct {
	ID            int       `gorm:"column:id"`
	ApplicationID int       `gorm:"column:application_id"`
	LogPath       string    `gorm:"column:log_path"`
	Added         time.Time `gorm:"column:added"`
}

type Logs []Log
