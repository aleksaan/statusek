package models

import (
	"errors"

	"github.com/aleksaan/statusek/config"
	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	ObjectName string
}

func (object *Object) TableName() string {
	// custom table name, this is default
	return config.Config.DBConfig.DbSchema + ".objects"
}

func (object *Object) GetObject(db *gorm.DB, objectName string) rc.ReturnCode {
	err := db.Where("object_name = ?", objectName).First(&object).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return rc.OBJECT_NAME_IS_NOT_FOUND
	}

	if err != nil {
		return rc.DATABASE_ERROR
	}

	return rc.SUCCESS
}
