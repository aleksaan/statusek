package logic

import (
	"fmt"

	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

var instanceInfo = &models.InstanceInfo{}
var statusInfo = &models.StatusInfo{}

func init() {

	//db initialization
	db = database.DB
	//db.LogMode(true)
}

func CheckInstanceIsFinished(instanceID int64) bool {

	//getting instance info
	instanceInfo.GetInstanceInfo(db, instanceID)

	//getting last statuses
	db.Raw("SELECT * FROM statuses.v_last_statuses WHERE object_id = ?", instanceInfo.Instance.ObjectID).Scan(&statusInfo.PrevStatuses)

	//checking previos statuses
	chk1 := CheckPreviosStatusesIsSet()

	return chk1
}

//CreateInstance - creates instance and gets its token
func CreateInstance(objectName string) (string, error) {
	objectID, err := GetObjectIDByName(objectName)
	if err != nil {
		return "", err
	}
	var instance = &models.Instance{ObjectID: objectID}
	db.Create(&instance)
	return instance.InstanceToken, nil
}

func GetObjectIDByName(objectName string) (int, error) {
	var object = &models.Object{}
	db.Debug().Where("object_name = ?", objectName).First(&object)
	if object.ObjectID > 0 {
		fmt.Printf("ObjectID: %d", object.ObjectID)
		return object.ObjectID, nil
	}
	return 0, fmt.Errorf("ERROR: Object name '" + objectName + "' has been not found in database")
}

func GetInstanceIDByToken(instanceToken string) (int64, error) {
	var instance = &models.Instance{}
	db.Debug().Where("instance_token::text = ?", instanceToken).First(&instance)
	if instance.InstanceID > 0 {
		fmt.Printf("InstanceID: %d", instance.InstanceID)
		return instance.InstanceID, nil
	}
	return 0, fmt.Errorf("ERROR: Instance token '" + instanceToken + "' has been not found in database")
}

func GetInstanceByToken(instanceToken string) (*models.Instance, error) {
	var instance = &models.Instance{}
	db.Debug().Where("instance_token::text = ?", instanceToken).First(&instance)
	if instance.InstanceID > 0 {
		fmt.Printf("InstanceID: %d", instance.InstanceID)
		return instance, nil
	}
	return nil, fmt.Errorf("ERROR: Instance token '" + instanceToken + "' has not been found in database")
}

func GetStatusIDByName(statusName string, objectID int) (int, error) {
	var status = &models.Status{}
	db.Debug().Where("status_name::text = ? and object_id = ?", statusName, objectID).First(&status)
	if status.StatusID > 0 {
		fmt.Printf("StatusID: %d", status.StatusID)
		return status.StatusID, nil
	}
	return 0, fmt.Errorf("ERROR: Status name '" + statusName + "' has been not found in database")
}

func SetStatus(instanceID int64, statusID int) error {

	tx := db.Begin()
	defer tx.Commit()

	//getting instance info
	instanceInfo.GetInstanceInfo(tx, instanceID)
	//instanceInfo.Print()

	//getting status info
	statusInfo.GetStatusInfo(tx, statusID)
	//statusInfo.Print()

	//checking status is according to instance
	chk0 := CheckStatusIsBelongsToInstance()

	//checking previos statuses
	chk1 := CheckPreviosStatusesIsSet()

	//checking next statuses
	chk2 := CheckNextStatusesIsNotSet()

	//cheking current status is not set yet
	chk3 := CheckCurrentStatusIsNotSet()

	fmt.Printf("\nchk1=%v, chk2=%v, chk3=%v\n", chk1, chk2, chk3)

	if chk0 && chk1 && chk2 && chk3 {
		event := &models.Event{StatusID: statusID, InstanceID: instanceID}
		err := tx.Create(&event).Error
		if err != nil {
			return err
		}
		event.Print()
	} else {
		return fmt.Errorf("ERROR: Status validation is not success")
	}

	return nil
}
