package models

import (
	"time"
)

type EventExtended struct {
	Status          Status
	EventMessage    string
	EventCreationDt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
