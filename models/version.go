package models

import (
	"github.com/aleksaan/statusek/config"
	"gorm.io/gorm"
)

//Version -
type Version struct {
	gorm.Model
	VersionNumber string
}

func (Version) TableName() string {
	return config.Config.DBConfig.DbSchema + ".version"
	//return "statusek.version"
}
