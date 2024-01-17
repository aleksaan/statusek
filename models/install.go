package models

import (
	"errors"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/logging"
	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/gorm"
)

var db = database.DB

func init() {
	UpdateDB()
}

func UpdateDB() rc.ReturnCode {
	//check existing of the version table
	var version = &Version{}
	checkTable := db.Migrator().HasTable(&Version{})

	//check version
	var checkVersion bool
	if checkTable {
		db.First(&version)
		checkVersion = version.VersionNumber == config.Config.DBVersion
	}

	var isVersionsAreDifferent = !checkTable || (checkTable && !checkVersion)

	//Если отсутствует таблица глобальных статусов - создаем ее
	checkGEvents := db.Migrator().HasTable(&GlobalEvent{})
	if !checkGEvents {
		db.AutoMigrate(&GlobalEvent{})
	}

	if config.Config.DBVersion == "v2023.06.14" && version.VersionNumber == "v2021.06.15_a" {
		logging.Info("Adding column 'message' to table 'events'...")
		db.Exec("ALTER TABLE " + config.Config.DBConfig.DbSchema + ".events ADD COLUMN IF NOT EXISTS message text;")
		logging.Info("Adding column 'message' to table 'events'... Done")
	}

	if isVersionsAreDifferent {
		logging.Info("Installed application version '%s' differs from current version '%s'", version.VersionNumber, config.Config.DBVersion)

		if !config.Config.DBConfig.DbUpdateIfOtherVersion {
			logging.Info("DB updating is canceled because parameter db_update_if_older_version=false")
			return rc.DB_IS_NOT_UPDATED
		}
	}

	if isVersionsAreDifferent && config.Config.DBConfig.DbUpdateIfOtherVersion {

		//creating DB objects
		logging.Info("Starting DB updating...")

		logging.Info("Dropping tables")
		db.Migrator().DropTable(&Event{}, &Workflow{}, &Instance{}, &Status{}, &Object{}, &Version{})

		logging.Info("Creating tables")
		db.AutoMigrate(&Version{}, &Object{}, &Instance{}, &Status{}, &Workflow{}, &Event{})

		//writing new version
		logging.Info("Writing new version number '%s'", config.Config.DBVersion)
		version := Version{VersionNumber: config.Config.DBVersion}
		CreateWrapper(db, &version)

		logging.Info("Creating default statuses models")
		CreatingDefaultModels(true)

		logging.Info("Starting DB updating...Done")
	} else {
		logging.Info("Current DB version is up to date")
	}

	return rc.SUCCESS

}

func CreatingDefaultModels(isUpdateAnyWhere bool) {
	logging.Info("Creating 2-POINT LINE TASK model...")
	//2-POINT LINE TASK

	//check existing of 2-point line task
	var objold = &Object{}
	err := db.Where("Object_name = ?", "2-POINT LINE TASK").First(&objold).Error

	if isUpdateAnyWhere || (!isUpdateAnyWhere && errors.Is(err, gorm.ErrRecordNotFound)) {

		if isUpdateAnyWhere && err == nil {
			//deleting old object
			db.Delete(&objold)
		}

		//creating new object
		obj := Object{ObjectName: "2-POINT LINE TASK"}
		CreateWrapper(db, &obj)
		st_start := Status{StatusName: "STARTED", ObjectID: obj.ID, StatusType: "MANDATORY"}
		CreateWrapper(db, &st_start)
		st_finish := Status{StatusName: "FINISHED", ObjectID: obj.ID, StatusType: "MANDATORY"}
		CreateWrapper(db, &st_finish)
		wkf := Workflow{StatusPrevID: st_start.ID, StatusNextID: st_finish.ID}
		CreateWrapper(db, &wkf)
	}
	logging.Info("Creating 2-POINT LINE TASK model...Done")
}
