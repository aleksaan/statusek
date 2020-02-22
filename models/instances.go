package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	u "github.com/aleksaan/scheduler/utils"
	"github.com/aleksaan/statusek/database"
)

type Instance struct {
	InstanceID         int64 `gorm:"primary_key;"`
	InstanceToken      string
	ObjectID           int
	InstanceCreationDt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (instance *Instance) BeforeCreate(scope *gorm.Scope) error {
	u := uuid.New().String()
	scope.SetColumn("InstanceToken", u)
	return nil
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
