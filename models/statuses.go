package models

import (
	u "github.com/aleksaan/scheduler/utils"
	"github.com/aleksaan/statusek/database"
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

func (status *Status) Create() map[string]interface{} {

	if err := database.DB.Create(status).Error; err != nil {
		errmsg := err.Error()
		resp := u.Message(false, errmsg)
		return resp
	}

	resp := u.Message(true, "success")
	resp["status"] = status
	return resp
}
