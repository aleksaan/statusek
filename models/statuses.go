package models

import (
	"github.com/aleksaan/statusek/config"
	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	ObjectID   uint
	Object     Object `json:"-"`
	StatusName string
	StatusDesc string
	StatusType string `gorm:"not null"`
}

func (status *Status) TableName() string {
	// custom table name, this is default
	return config.Config.DBConfig.DbSchema + ".statuses"
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
	tx.Where("id = ?", statusId).First(&status)
	return rc.SUCCESS
}
