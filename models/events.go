package models

import (
	"time"

	u "github.com/aleksaan/scheduler/utils"
	"github.com/aleksaan/statusek/database"
)

type Event struct {
	EventID         int64 `gorm:"primary_key;`
	InstanceID      int
	StatusID        int
	EventCreationDt *time.Time
}

func (event *Event) TableName() string {
	// custom table name, this is default
	return "statuses.events"
}

func (event *Event) Create() map[string]interface{} {

	if err := database.DB.Create(event).Error; err != nil {
		errmsg := err.Error()
		resp := u.Message(false, errmsg)
		return resp
	}

	resp := u.Message(true, "success")
	resp["event"] = event
	return resp
}
