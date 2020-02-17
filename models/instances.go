package models

import (
	"time"

	u "github.com/aleksaan/scheduler/utils"
	"github.com/aleksaan/statusek/database"
)

type Instance struct {
	InstanceID         int `gorm:"primary_key;"`
	InstanceToken      string
	ObjectID           int
	InstanceCreationDt *time.Time
}

func (instance *Instance) TableName() string {
	// custom table name, this is default
	return "statuses.instances"
}

func (instance *Instance) Create() map[string]interface{} {

	if err := database.DB.Create(instance).Error; err != nil {
		errmsg := err.Error()
		resp := u.Message(false, errmsg)
		return resp
	}

	resp := u.Message(true, "success")
	resp["instance"] = instance
	return resp
}
