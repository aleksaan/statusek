package models

import (
	"gorm.io/gorm"
)

type GlobalEvent struct {
	gorm.Model
	EventName    string
	EventMessage string
	StatusType   string `gorm:"not null"`
}
