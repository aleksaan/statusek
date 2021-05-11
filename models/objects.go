package models

import (
	"errors"

	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/jinzhu/gorm"
)

type Object struct {
	gorm.Model
	ObjectName string
	Instance   []Instance `gorm:"ForeignKey:InstanceID"`
	Status     []Status   `gorm:"ForeignKey:StatusID"`
}

func (object *Object) TableName() string {
	// custom table name, this is default
	return "statusek.objects"
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

// func (object *Object) Create() map[string]interface{} {

// 	if err := database.DB.Create(object).Error; err != nil {
// 		errmsg := err.Error()
// 		resp := u.Message(false, errmsg)
// 		return resp
// 	}

// 	resp := u.Message(true, "success")
// 	resp["object"] = object
// 	return resp
// }
