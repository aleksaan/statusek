package models

import (
	"github.com/aleksaan/statusek/logging"
	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/gorm"
)

//CreateWrapper - wrapper for database create function (gets returncode and write database error to log)
func CreateWrapper(tx *gorm.DB, v interface{}) rc.ReturnCode {

	err := tx.Create(v).Error
	if err != nil {
		logging.Error("%s", err.Error())
		return rc.DATABASE_ERROR
	}
	return rc.SUCCESS
}
