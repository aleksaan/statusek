package models

import (
	"fmt"

	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
	"gorm.io/gorm"
)

type InstanceInfo struct {
	Instance Instance
	Events   []Event
	Statuses []Status
}

func (instanceInfo *InstanceInfo) GetInstanceInfo(db *gorm.DB, instanceToken string, isForUpdate bool) rc.ReturnCode {
	rc0 := instanceInfo.Instance.GetInstance(db, instanceToken, isForUpdate)

	if rc0 != rc.SUCCESS {
		return rc0
	}

	instanceInfo.RefreshEvents(db)
	instanceInfo.RefreshStatuses(db)

	return rc.SUCCESS
}

func (instanceInfo *InstanceInfo) RefreshEvents(db *gorm.DB) rc.ReturnCode {
	db.Where("instance_id = ?", instanceInfo.Instance.ID).Find(&instanceInfo.Events)
	return rc.SUCCESS
}

func (instanceInfo *InstanceInfo) RefreshStatuses(db *gorm.DB) rc.ReturnCode {
	db.Where("object_id = ?", instanceInfo.Instance.ObjectID).Find(&instanceInfo.Statuses)
	return rc.SUCCESS
}

func (instanceInfo *InstanceInfo) ToString() string {
	return utils.ToString(&instanceInfo)
}

func (instanceInfo *InstanceInfo) Print() {
	fmt.Printf("\n-----------------------Instance-----------------------\n")
	fmt.Printf("%s\n\n", instanceInfo.ToString())
}
