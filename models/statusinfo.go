package models

import (
	"fmt"

	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
	"github.com/jinzhu/gorm"
)

type StatusInfo struct {
	Status       Status
	PrevStatuses []Status
	NextStatuses []Status
}

//GetStatusInfo - get status info by statusName & objectID
func (statusInfo *StatusInfo) GetStatusInfo(tx *gorm.DB, statusName string, objectID int) rc.ReturnCode {

	rc0 := statusInfo.Status.GetStatus(tx, statusName, objectID)
	if rc0 != rc.SUCCESS {
		return rc0
	}

	var workflows []Workflow
	tx.Where("status_id_next = ?", statusInfo.Status.StatusID).Find(&workflows)
	var statusesIDPrev []int
	for _, w := range workflows {
		statusesIDPrev = append(statusesIDPrev, w.StatusIDPrev)
	}
	tx.Where("status_id in (?)", statusesIDPrev).Find(&statusInfo.PrevStatuses)

	tx.Where("status_id_prev = ?", statusInfo.Status.StatusID).Find(&workflows)
	var statusesIDNext []int
	for _, w := range workflows {
		statusesIDNext = append(statusesIDNext, w.StatusIDNext)
	}
	tx.Where("status_id in (?)", statusesIDNext).Find(&statusInfo.NextStatuses)

	return rc.SUCCESS
}

func (statusInfo *StatusInfo) ToString() string {
	return utils.ToString(&statusInfo)
}

func (statusInfo *StatusInfo) Print() {
	fmt.Printf("\n-----------------------Status-----------------------\n")
	fmt.Printf("%s\n\n", statusInfo.ToString())
}
