package models

import (
	"time"
)

type EventExtended struct {
	Status          Status
	EventCreationDt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
