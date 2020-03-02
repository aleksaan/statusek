package models

import (
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/jinzhu/gorm"
)

type Status struct {
	ObjectID          int
	StatusID          int `gorm:"primary_key;`
	StatusName        string
	StatusIsMandatory bool
	StatusDesc        string
}

func (status *Status) TableName() string {
	// custom table name, this is default
	return "statuses.statuses"
}

func (status *Status) GetStatus(tx *gorm.DB, statusName string, objectID int) rc.ReturnCode {
	tx.Where("status_name = ? and object_id = ?", statusName, objectID).First(&status)

	if status.StatusID > 0 {
		//fmt.Printf("StatusID: %d", status.StatusID)
		return rc.SUCCESS
	}
	return rc.STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT
}

// func (status *Status) Create() map[string]interface{} {

// 	if err := database.DB.Create(status).Error; err != nil {
// 		errmsg := err.Error()
// 		resp := u.Message(false, errmsg)
// 		return resp
// 	}

// 	resp := u.Message(true, "success")
// 	resp["status"] = status
// 	return resp
// }
