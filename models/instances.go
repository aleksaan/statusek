package models

import (
	"errors"

	"github.com/aleksaan/statusek/config"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Instance struct {
	gorm.Model
	InstanceToken                 string
	InstanceTimeout               int
	ObjectID                      uint
	Object                        Object
	InstanceIsFinished            bool
	InstanceIsFinishedDescription string
}

func (instance *Instance) BeforeCreate(tx *gorm.DB) (err error) {
	instance.InstanceToken = uuid.New().String()
	return
}

func (instance *Instance) TableName() string {
	// custom table name, this is default
	return config.Config.DBConfig.DbSchema + ".instances"
}

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
