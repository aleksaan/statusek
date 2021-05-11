package models

import (
	"github.com/jinzhu/gorm"
)

//Version -
type Version struct {
	gorm.Model
	VersionNumber string
}

//TableName -
func (version *Version) TableName() string {
	// custom table name, this is default
	return "statusek.version"
}
