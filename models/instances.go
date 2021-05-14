package models

import (
	"errors"

	"github.com/aleksaan/statusek/database"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Instance struct {
	gorm.Model
	InstanceToken                 string
	InstanceTimeout               int
	ObjectID                      uint
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
	return database.ConnectionSettings.DbSchema + ".instances"
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
	err := db.Set("gorm:query_option", option).Where("instance_token = ?", instanceToken).First(&instance).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rc.INSTANCE_TOKEN_IS_NOT_FOUND
	}

	if err != nil {
		return rc.DATABASE_ERROR
	}

	return rc.SUCCESS

}

func (instance *Instance) FinishInstance(db *gorm.DB, instanceIsFinishedDescription string) rc.ReturnCode {
	instance.InstanceIsFinished = true
	instance.InstanceIsFinishedDescription = instanceIsFinishedDescription
	err := db.Save(&instance)
	if err != nil {
		return rc.DATABASE_ERROR
	}
	return rc.SUCCESS
}
