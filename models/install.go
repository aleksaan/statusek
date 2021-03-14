package models

import (
	"github.com/aleksaan/statusek/database"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = database.DB
}

func UpdateDB(currentVersion string) {
	//check existing of the version table
	var checkTable = db.HasTable(&Version{})
	var checkVersion bool

	//check version
	if checkTable {
		var version = &Version{}
		db.First(&version)
		checkVersion = version.VersionNumber == currentVersion
	}

	if !checkTable || (checkTable && !checkVersion) {

		//creating DB objects
		db.LogMode(true)
		db.DropTable(&Event{}, &Workflow{}, &Instance{}, &Status{}, &Object{}, &Version{})
		db.AutoMigrate(&Version{}, &Object{}, &Instance{}, &Status{}, &Workflow{}, &Event{})
		db.Model(&Instance{}).AddForeignKey("object_id", "objects(id)", "CASCADE", "CASCADE")
		db.Model(&Status{}).AddForeignKey("object_id", "objects(id)", "CASCADE", "CASCADE")
		db.Model(&Workflow{}).AddForeignKey("status_id_prev", "statuses(id)", "CASCADE", "CASCADE")
		db.Model(&Workflow{}).AddForeignKey("status_id_next", "statuses(id)", "CASCADE", "CASCADE")
		db.Model(&Event{}).AddForeignKey("instance_id", "instances(id)", "CASCADE", "CASCADE")
		db.Model(&Event{}).AddForeignKey("status_id", "statuses(id)", "CASCADE", "CASCADE")

		//writing new version
		version := Version{VersionNumber: currentVersion}
		result := db.Create(&version)
		if result.Error != nil {
			//error handler
		}

	}

}
