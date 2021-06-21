package models

import (
	"fmt"
	"time"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/logging"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	InstanceID      uint
	Instance        Instance
	StatusID        uint
	Status          Status
	EventCreationDt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (event *Event) TableName() string {
	// custom table name, this is default
	return config.Config.DBConfig.DbSchema + ".events"
}

func (event *Event) ToString() string {
	return utils.ToString(&event)
}

func (event *Event) Print() {
	fmt.Printf("\n-----------------------Event-----------------------\n")
	fmt.Printf("%s\n\n", event.ToString())
}

func (event *Event) Create(tx *gorm.DB) rc.ReturnCode {
	res := tx.Create(&event)
	if res.Error != nil {
		logging.Error(res.Error.Error())
		return rc.DATABASE_ERROR
	}
	return rc.SUCCESS
}
