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

func finishInstanceIfTimeout(instanceInfo *models.InstanceInfo) {
	if instanceInfo.Instance.InstanceIsFinished == true {
		return
	}
	tx := db.Begin()
	defer tx.Commit()
	chk4, _ := checkInstanceIsNotTimeout(instanceInfo)
	if !chk4 {
		instanceInfo.Instance.FinishInstance(tx, "TIMEOUT")
	}
}

func GetEvents(instanceToken string) ([]models.EventExtended, rc.ReturnCode) {
	var events []models.EventExtended

	var instanceInfo = &models.InstanceInfo{}

	//getting instance info
	rc5 := instanceInfo.GetInstanceInfo(db, instanceToken, false)
	if rc5 != rc.SUCCESS {
		return events, rc5
	}
	finishInstanceIfTimeout(instanceInfo)

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

// CheckInstanceIsFinished - check for finishing
func CheckInstanceIsFinished(instanceToken string) (bool, rc.ReturnCode) {
	var instanceInfo = &models.InstanceInfo{}
	tx := db.Begin()

	//getting instance info (FOR UPDATE MODE)
	rc5 := instanceInfo.GetInstanceInfo(tx, instanceToken, true)
	if rc5 != rc.SUCCESS {
		tx.Rollback()
		return false, rc5
	}

	finishInstanceIfTimeout(instanceInfo)
	tx.Commit()
	return checkInstanceIsFinished(instanceInfo)
}

// SetStatus - set status of instance
func SetStatus(instanceToken string, statusName string) rc.ReturnCode {

	tx := db.Begin()

	var instanceInfo = &models.InstanceInfo{}
	var statusInfo = &models.StatusInfo{}

	//getting instance info (FOR UPDATE MODE)
	rc5 := instanceInfo.GetInstanceInfo(tx, instanceToken, true)
	if rc5 != rc.SUCCESS {
		return rc5
	}

	finishInstanceIfTimeout(instanceInfo)
	if chk7, rc7 := checkInstanceIsFinished(instanceInfo); chk7 == true {
		return rc7
	}

	//getting status info
	rc6 := statusInfo.GetStatusInfo(tx, statusName, instanceInfo.Instance.ObjectID)
	if rc6 != rc.SUCCESS {
		tx.Rollback()
		return rc6
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
	instanceInfo.Events = append(instanceInfo.Events, *event)

	//finish instance if stop-status got
	if statusInfo.Status.StatusType == "STOP-STATUS" {
		instanceInfo.Instance.FinishInstance(tx, "STOP_STATUS_IS_SET_2")
	}

	//finish instance if all mandatory statuses is set
	if r, _ := checkAllMandatoryStatusesAreSet(instanceInfo); r == true {
		instanceInfo.Instance.FinishInstance(tx, "ALL_MANDATORY_STATUSES_ARE_SET")
	}

	tx.Commit()
	return rc.SUCCESS
}
