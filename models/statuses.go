package models

import (
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/jinzhu/gorm"
)

type Status struct {
	gorm.Model
	ObjectID   uint
	Object     Object
	StatusName string
	StatusDesc string
	StatusType string
	Workflow   []Workflow `gorm:"ForeignKey:WorkflowID"`
	Event      []Event    `gorm:"ForeignKey:EventID"`
}

func (status *Status) TableName() string {
	// custom table name, this is default
	return "statusek.statuses"
}

func (status *Status) GetStatus(tx *gorm.DB, statusName string, objectID uint) rc.ReturnCode {
	tx.Where("status_name = ? and object_id = ?", statusName, objectID).First(&status)

	if status.ID > 0 {
		//fmt.Printf("StatusID: %d", status.StatusID)
		return rc.SUCCESS
	}
	return rc.STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT
}

func (status *Status) GetStatusById(tx *gorm.DB, statusId uint) rc.ReturnCode {
	tx.Where("status_id = ?", statusId).First(&status)

	if status.ID > 0 {
		//fmt.Printf("StatusID: %d", status.StatusID)
		return rc.SUCCESS
	}
	return rc.STATUS_ID_IS_NOT_FOUND
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
