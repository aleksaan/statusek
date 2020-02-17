package models

import (
	u "github.com/aleksaan/scheduler/utils"
	"github.com/aleksaan/statusek/database"
)

//Workflow -
type Workflow struct {
	WorkflowID   int `gorm:"primary_key;`
	StatusIDPrev int
	StatusIDNext int
}

//TableName -
func (workflow *Workflow) TableName() string {
	// custom table name, this is default
	return "statuses.workflows"
}

//Create -
func (workflow *Workflow) Create() map[string]interface{} {

	if err := database.DB.Create(workflow).Error; err != nil {
		errmsg := err.Error()
		resp1 := u.Message(false, errmsg)
		return resp1
	}

	resp := u.Message(true, "success")
	resp["workflow"] = workflow
	return resp
}

// object_id
// status_id_prev
// status_name_prev
// status_is_mandatory_prev
// status_id
// status_name
// status_is_mandatory
// status_level
