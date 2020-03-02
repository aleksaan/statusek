package models

import (
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/jinzhu/gorm"
)

type Object struct {
	ObjectID   int `gorm:"primary_key;"`
	ObjectName string
}

func (object *Object) TableName() string {
	// custom table name, this is default
	return "statuses.objects"
}

func (object *Object) GetObject(db *gorm.DB, objectName string) rc.ReturnCode {
	db.Where("object_name = ?", objectName).First(&object)
	if object.ObjectID > 0 {
		return rc.SUCCESS
	}
	return rc.OBJECT_NAME_IS_NOT_FOUND
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
