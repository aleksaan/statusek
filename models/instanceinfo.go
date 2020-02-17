package models

import (
	"fmt"

	"github.com/aleksaan/statusek/utils"
	"github.com/jinzhu/gorm"
)

type InstanceInfo struct {
	Instance Instance
	Events   []Event
}

func (instanceInfo *InstanceInfo) GetInstanceInfo(db *gorm.DB, instanceID int64) {
	db.Debug().Set("gorm:query_option", "FOR UPDATE").Where("instance_id = (?)", instanceID).First(&instanceInfo.Instance)
	db.Where("instance_id = ?", instanceID).Find(&instanceInfo.Events)
}

func (instanceInfo *InstanceInfo) ToString() string {
	return utils.ToString(&instanceInfo)
}

func (instanceInfo *InstanceInfo) Print() {
	fmt.Printf("\n-----------------------Instance-----------------------\n")
	fmt.Printf("%s\n\n", instanceInfo.ToString())
}
