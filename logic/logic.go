package logic

import (
	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {

	db = database.DB
}

//------------------------ADVANCED FUNCTIONS----------------------------------

// CreateInstance - creates instance of object and gets its token

func CreateInstance(objectName string, instanceTimeout int) (string, rc.ReturnCode) {
	object := &models.Object{}
	rc0 := object.GetObject(db, objectName)
	if rc0 != rc.SUCCESS {
		return "", rc0
	}
	var instance = &models.Instance{ObjectID: object.ObjectID, InstanceTimeout: instanceTimeout}
	db.Create(&instance)
	return instance.InstanceToken, rc.SUCCESS
}

func CheckInstanceIsFinished(instanceToken string) (bool, rc.ReturnCode) {
	//getting instance info
	var instanceInfo = &models.InstanceInfo{}
	rc5 := instanceInfo.GetInstanceInfo(db, instanceToken)
	if rc5 != rc.SUCCESS {
		return false, rc5
	}

	return checkInstanceIsFinished(instanceInfo)
}

// SetStatus - set status of instance

func SetStatus(instanceToken string, statusName string) rc.ReturnCode {

	tx := db.Begin()

	var instanceInfo = &models.InstanceInfo{}
	var statusInfo = &models.StatusInfo{}

	//getting instance info
	rc5 := instanceInfo.GetInstanceInfo(tx, instanceToken)
	if rc5 != rc.SUCCESS {
		return rc5
	}

	//getting status info
	rc6 := statusInfo.GetStatusInfo(tx, statusName, instanceInfo.Instance.ObjectID)
	if rc6 != rc.SUCCESS {
		tx.Rollback()
		return rc6
	}
	//statusInfo.Print()

	//checking instance is timed out
	chk4, rc4 := checkInstanceIsNotTimeout(instanceInfo)
	if !chk4 {
		tx.Rollback()
		return rc4
	}

	//checking status is according to instance
	chk0, rc0 := checkStatusIsBelongsToInstance(instanceInfo, statusInfo)
	if !chk0 {
		tx.Rollback()
		return rc0
	}

	//checking previos statuses
	chk1, rc1 := checkPreviosStatusesIsSet(instanceInfo, statusInfo)
	if !chk1 {
		tx.Rollback()
		return rc1
	}

	//checking next statuses
	chk2, rc2 := checkNextStatusesIsNotSet(instanceInfo, statusInfo)
	if !chk2 {
		tx.Rollback()
		return rc2
	}

	//cheking current status is not set yet
	chk3, rc3 := checkCurrentStatusIsNotSet(instanceInfo, statusInfo)
	if !chk3 {
		tx.Rollback()
		return rc3
	}

	event := &models.Event{StatusID: statusInfo.Status.StatusID, InstanceID: instanceInfo.Instance.InstanceID}
	err := tx.Create(&event).Error
	if err != nil {
		tx.Rollback()
		return rc.SET_STATUS_DB_ERROR
	}

	tx.Commit()
	return rc.SUCCESS
}
