package models

import (
	"fmt"

	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
	"github.com/jinzhu/gorm"
)

type InstanceInfo struct {
	Instance Instance
	Events   []Event
}

func (instanceInfo *InstanceInfo) GetInstanceInfo(db *gorm.DB, instanceToken string) rc.ReturnCode {
	rc0 := instanceInfo.Instance.GetInstance(db, instanceToken, true)

	if rc0 != rc.SUCCESS {
		return rc0
	}
	db.Where("instance_id = ?", instanceInfo.Instance.InstanceID).Find(&instanceInfo.Events)
	return rc.SUCCESS
}

func (instanceInfo *InstanceInfo) ToString() string {
	return utils.ToString(&instanceInfo)
}

func (instanceInfo *InstanceInfo) Print() {
	fmt.Printf("\n-----------------------Instance-----------------------\n")
	fmt.Printf("%s\n\n", instanceInfo.ToString())
}
