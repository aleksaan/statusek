package models

import (
	"fmt"

	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
	"gorm.io/gorm"
)

type StatusInfo struct {
	Status       Status
	PrevStatuses []Status
	NextStatuses []Status
}

//GetStatusInfo - get status info by statusName & objectID
func (statusInfo *StatusInfo) GetStatusInfo(tx *gorm.DB, statusName string, objectID uint) rc.ReturnCode {

	rc0 := statusInfo.Status.GetStatus(tx, statusName, objectID)
	if rc0 != rc.SUCCESS {
		return rc0
	}

	var workflows []Workflow
	tx.Where("status_next_id = ?", statusInfo.Status.ID).Find(&workflows)
	var statusesIDPrev []int
	for _, w := range workflows {
		statusesIDPrev = append(statusesIDPrev, int(w.StatusPrevID))
	}
	tx.Where("id in (?)", statusesIDPrev).Find(&statusInfo.PrevStatuses)

	tx.Where("status_prev_id = ?", statusInfo.Status.ID).Find(&workflows)
	var statusesIDNext []int
	for _, w := range workflows {
		statusesIDNext = append(statusesIDNext, int(w.StatusNextID))
	}
	tx.Where("id in (?)", statusesIDNext).Find(&statusInfo.NextStatuses)

	return rc.SUCCESS
}

func (statusInfo *StatusInfo) ToString() string {
	return utils.ToString(&statusInfo)
}

func (statusInfo *StatusInfo) Print() {
	fmt.Printf("\n-----------------------Status-----------------------\n")
	fmt.Printf("%s\n\n", statusInfo.ToString())
}
