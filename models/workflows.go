package models

import (
	"github.com/aleksaan/statusek/config"
	"gorm.io/gorm"
)

//Workflow -
type Workflow struct {
	gorm.Model
	StatusPrevID uint
	StatusPrev   Status
	StatusNextID uint
	StatusNext   Status
}

//TableName -
func (workflow *Workflow) TableName() string {
	// custom table name, this is default
	return config.Config.DBConfig.DbSchema + ".workflows"
}

// object_id
// status_id_prev
// status_name_prev
// status_is_mandatory_prev
// status_id
// status_name
// status_is_mandatory
// status_level
