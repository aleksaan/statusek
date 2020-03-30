package models

import (
	"time"

	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Instance struct {
	InstanceID                    int64 `gorm:"primary_key;"`
	InstanceToken                 string
	InstanceTimeout               int
	ObjectID                      int
	InstanceCreationDt            *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	InstanceIsFinished            bool
	InstanceIsFinishedDescription string
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

// func (instance *Instance) Create() map[string]interface{} {

// 	if err := database.DB.Create(instance).Error; err != nil {
// 		errmsg := err.Error()
// 		resp := u.Message(false, errmsg)
// 		return resp
// //	}

// 	resp := u.Message(true, "success")
// 	resp["instance"] = instance
// 	return resp
// }

func (instance *Instance) GetInstance(db *gorm.DB, instanceToken string, isForUpdate bool) rc.ReturnCode {
	option := ""
	if isForUpdate {
		option = "FOR UPDATE"
	}
	db.Set("gorm:query_option", option).Where("instance_token = ?", instanceToken).First(&instance)
	if instance.InstanceID > 0 {
		//fmt.Printf("InstanceID: %d", instance.InstanceID)
		return rc.SUCCESS
	}
	return rc.INSTANCE_TOKEN_IS_NOT_FOUND
}

func (instance *Instance) FinishInstance(db *gorm.DB, instanceIsFinishedDescription string) rc.ReturnCode {
	instance.InstanceIsFinished = true
	instance.InstanceIsFinishedDescription = instanceIsFinishedDescription
	db.Save(&instance)
	return rc.SUCCESS
}
