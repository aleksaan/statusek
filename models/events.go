package models

import (
	"fmt"
	"time"

	"github.com/aleksaan/statusek/utils"
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	InstanceID      uint
	StatusID        uint
	EventCreationDt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (event *Event) TableName() string {
	// custom table name, this is default
	return "statusek.events"
}

// func (event *Event) Create() map[string]interface{} {

// 	if err := database.DB.Create(event).Error; err != nil {
// 		errmsg := err.Error()
// 		resp := u.Message(false, errmsg)
// 		return resp
// 	}

// 	resp := u.Message(true, "success")
// 	resp["event"] = event
// 	return resp
// }

func (event *Event) ToString() string {
	return utils.ToString(&event)
}

func (event *Event) Print() {
	fmt.Printf("\n-----------------------Event-----------------------\n")
	fmt.Printf("%s\n\n", event.ToString())
}
