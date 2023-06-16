package logic

import (
	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/gorm"
)

var db = database.DB

// CreateInstance - creates instance of object and gets its token
func CreateInstance(objectName string, instanceTimeout int) (string, rc.ReturnCode) {
	object := &models.Object{}
	rc0 := object.GetObject(db, objectName)
	if rc0 != rc.SUCCESS {
		return "", rc0
	}
	var instance = &models.Instance{ObjectID: object.ID, InstanceTimeout: instanceTimeout}
	rc1 := models.CreateWrapper(db, instance)
	if rc1 != rc.SUCCESS {
		return "", rc1
	}

	return instance.InstanceToken, rc.SUCCESS
}

func finishInstanceIfTimeout(tx *gorm.DB, instanceInfo *models.InstanceInfo) {
	if instanceInfo.Instance.InstanceIsFinished {
		return
	}
	rc0 := checkInstanceIsTimeout(instanceInfo)
	if rc0 == rc.INSTANCE_IS_TIMEOUT {
		instanceInfo.Instance.FinishInstance(tx, "TIMEOUT")
	}
}

func GetEvents(instanceToken string) ([]models.EventExtended, rc.ReturnCode) {
	var events []models.EventExtended

	var instanceInfo = &models.InstanceInfo{}
	tx := db.Begin()
	defer tx.Commit()

	//getting instance info (FOR UPDATE MODE)
	rc5 := instanceInfo.GetInstanceInfo(tx, instanceToken, false)
	if rc5 != rc.SUCCESS {
		return events, rc5
	}
	finishInstanceIfTimeout(tx, instanceInfo)

	var status = &models.Status{}

	for _, e := range instanceInfo.Events {
		rc0 := status.GetStatusById(db, e.StatusID)
		if rc0 != rc.SUCCESS {
			return nil, rc0
		}
		event := &models.EventExtended{EventCreationDt: e.EventCreationDt}
		event.Status.GetStatusById(db, e.StatusID)
		events = append(events, *event)
	}

	return events, rc.SUCCESS
}

// CheckStatusIsSet - check certain status is set
func CheckStatusIsSet(instanceToken string, statusName string) (bool, rc.ReturnCode) {
	var instanceInfo = &models.InstanceInfo{}
	tx := db.Begin()
	defer tx.Commit()

	//getting instance info (FOR UPDATE MODE)
	rc5 := instanceInfo.GetInstanceInfo(tx, instanceToken, true)
	if rc5 != rc.SUCCESS {
		return false, rc5
	}
	finishInstanceIfTimeout(tx, instanceInfo)

	var status = &models.Status{}
	rc1 := status.GetStatus(tx, statusName, instanceInfo.Instance.ObjectID)
	if rc1 != rc.SUCCESS {
		return false, rc1
	}

	//looking for status is set
	rc0 := checkStatusIsSet(instanceInfo, status)
	if rc0 == rc.STATUS_IS_SET {
		return true, rc0
	}
	return false, rc0
}

// GetInstanceInfo - check for finishing
func GetInstanceInfo(instanceToken string) (bool, rc.ReturnCode, *models.InstanceInfo) {
	var instanceInfo = &models.InstanceInfo{}
	tx := db.Begin()
	defer tx.Commit()
	rc0 := instanceInfo.GetInstanceInfo(tx, instanceToken, true)
	if rc0 != rc.SUCCESS {
		return false, rc0, instanceInfo
	}
	finishInstanceIfTimeout(tx, instanceInfo)
	return true, rc0, instanceInfo
}

// SetStatus - set status of instance
func SetStatus(instanceToken string, statusName string, statusMessage string) rc.ReturnCode {
	//logging.Error("[%s] [%s] Setting status '%s' for instance '%s'... Error: '%s'", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken, rc3message)

	tx := db.Begin()

	//tx := db.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	var instanceInfo = &models.InstanceInfo{}
	var statusInfo = &models.StatusInfo{}

	rc0 := checkStatusIsReadyToSet(tx, instanceInfo, statusInfo, instanceToken, statusName)
	if rc0 != rc.SUCCESS {
		return rc0
	}

	event := &models.Event{StatusID: statusInfo.Status.ID, InstanceID: instanceInfo.Instance.ID, Message: statusMessage}

	rc8 := models.CreateWrapper(tx, event)
	if rc8 != rc.SUCCESS {
		return rc8
	}

	instanceInfo.Events = append(instanceInfo.Events, *event)

	//finish instance if stop-status got
	if statusInfo.Status.StatusType == "STOP-STATUS" {
		instanceInfo.Instance.FinishInstance(tx, "STOP_STATUS_IS_SET")
	}

	//finish instance if all mandatory statuses is set
	if r, _ := checkAllMandatoryStatusesAreSet(instanceInfo); r {
		instanceInfo.Instance.FinishInstance(tx, "ALL_MANDATORY_STATUSES_ARE_SET")
	}

	return rc.SUCCESS
}

func CheckStatusIsReadyToSet(instanceToken string, statusName string) rc.ReturnCode {
	tx := db.Begin()
	defer tx.Commit()

	var instanceInfo = &models.InstanceInfo{}
	var statusInfo = &models.StatusInfo{}

	return checkStatusIsReadyToSet(tx, instanceInfo, statusInfo, instanceToken, statusName)
}

func CloneProcess(instanceToken string) (string, rc.ReturnCode) {
	var instanceInfo = &models.InstanceInfo{}
	tx := db.Begin()
	defer tx.Commit()
	rc0 := instanceInfo.GetInstanceInfo(tx, instanceToken, true)
	if rc0 != rc.SUCCESS {
		return "", rc0
	}

	newToken, rc1 := CreateInstance(instanceInfo.Object.ObjectName, instanceInfo.Instance.InstanceTimeout)
	if rc1 != rc.SUCCESS {
		return "", rc1
	}

	for _, v := range instanceInfo.Events {
		var status = &models.Status{}
		db.First(&status, v.StatusID)
		if status.StatusType != "STOP-STATUS" {
			rc2 := SetStatus(newToken, status.StatusName, v.Message)
			if rc2 != rc.SUCCESS {
				return "", rc2
			}
		}
	}

	return newToken, rc.SUCCESS
}
