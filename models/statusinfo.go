package models

import (
	"fmt"

	"github.com/aleksaan/statusek/utils"
	"github.com/jinzhu/gorm"
)

type StatusInfo struct {
	Status       Status
	PrevStatuses []Status
	NextStatuses []Status
}

func (statusInfo *StatusInfo) GetStatusInfo(db *gorm.DB, statusID int) {
	db.Where("status_id = ?", statusID).Find(&statusInfo.Status)

	var workflows []Workflow
	db.Where("status_id_next = ?", statusID).Find(&workflows)
	var statusesIDPrev []int
	for _, w := range workflows {
		statusesIDPrev = append(statusesIDPrev, w.StatusIDPrev)
	}
	db.Where("status_id in (?)", statusesIDPrev).Find(&statusInfo.PrevStatuses)

	db.Where("status_id_prev = ?", statusID).Find(&workflows)
	var statusesIDNext []int
	for _, w := range workflows {
		statusesIDNext = append(statusesIDNext, w.StatusIDNext)
	}
	db.Where("status_id in (?)", statusesIDNext).Find(&statusInfo.NextStatuses)
}

func (statusInfo *StatusInfo) ToString() string {
	return utils.ToString(&statusInfo)
}

func (statusInfo *StatusInfo) Print() {
	fmt.Printf("\n-----------------------Status-----------------------\n")
	fmt.Printf("%s\n\n", statusInfo.ToString())
}
