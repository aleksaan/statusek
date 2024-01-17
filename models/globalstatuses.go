package models

import (
	"github.com/aleksaan/statusek/config"
	"gorm.io/gorm"
)

type GlobalEvent struct {
	gorm.Model
	EventName string `gorm:"not null"`
}

func (event *GlobalEvent) TableName() string {
	// custom table name, this is default
	return config.Config.DBConfig.DbSchema + ".globalevents"
}
