package logic

import (
	"time"

	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/gorm"
)

func checkStatusIsBelongsToInstance(instanceInfo *models.InstanceInfo, statusInfo *models.StatusInfo) (bool, rc.ReturnCode) {
	if instanceInfo.Instance.ObjectID == statusInfo.Status.ObjectID {
		return true, rc.SUCCESS
	}
	return false, rc.STATUS_IS_NOT_ACCORDING_TO_INSTANCE
}

func checkAllMandatoryStatusesAreSet(instanceInfo *models.InstanceInfo) (bool, rc.ReturnCode) {

	var countMandatoryOverall = 0
	var countMandatoryIsSet = 0
	for _, s := range instanceInfo.Statuses {
		if s.StatusType == "MANDATORY" {
			countMandatoryOverall++
			for _, e := range instanceInfo.Events {
				if s.ID == e.StatusID {
					countMandatoryIsSet++
				}
			}

		}
	}

	if countMandatoryOverall == countMandatoryIsSet {
		return true, rc.ALL_MANDATORY_ARE_SET
	}

	return false, rc.NOT_ALL_MANDATORY_ARE_SET
}

// CheckInstanceIsFinished - checks if instance finished or not
// Finished is if all of mandatory statuses of last level is set or if no one mandatory
// then at least one of optional statuses is set

func checkInstanceIsFinished(instanceInfo *models.InstanceInfo) (bool, rc.ReturnCode) {

	if instanceInfo.Instance.InstanceIsFinished {
		return true, rc.INSTANCE_IS_FINISHED
	}

	return false, rc.INSTANCE_IS_NOT_FINISHED
}

func checkInstanceIsNotTimeout(instanceInfo *models.InstanceInfo) (bool, rc.ReturnCode) {
	t1 := time.Now()
	t2 := instanceInfo.Instance.CreatedAt
	//fmt.Printf("\n\nTime 1: %s\nTime 2: %s\n\n", t1.Format(time.RFC3339), t2.Format(time.RFC3339))
	diff := t1.Sub(t2).Seconds()
	if diff < float64(instanceInfo.Instance.InstanceTimeout) {
		return true, rc.SUCCESS
	}

	return false, rc.INSTANCE_IS_IN_TIMEOUT
}

func checkPreviosStatusesIsSet(instanceInfo *models.InstanceInfo, statusInfo *models.StatusInfo) (bool, rc.ReturnCode) {
	var countPrevMandatory int
	var countPrevMandatoryIsSet int
	var countPrevOptional int
	var countPrevOptionalIsSet int
	for _, s := range statusInfo.PrevStatuses {
		if s.StatusType == "MANDATORY" {
			countPrevMandatory++
			for _, e := range instanceInfo.Events {
				if e.StatusID == s.ID {
					countPrevMandatoryIsSet++
					break
				}
			}
		} else {
			countPrevOptional++
			for _, e := range instanceInfo.Events {
				if e.StatusID == s.ID {
					countPrevOptionalIsSet++
					break
				}
			}
		}
	}

	if countPrevMandatory > countPrevMandatoryIsSet {
		return false, rc.NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET
		//"Не все обязательные статусы предыдущего уровня установлены"
	}

	if (countPrevMandatory == 0) && (countPrevOptional > 0) && (countPrevOptionalIsSet == 0) {
		return false, rc.NO_ONE_PREVIOS_OPTIONAL_STATUSES_ARE_SET
		//"Не установлен ни один опциональный статус предыдущего уровня"
	}

	return true, rc.SUCCESS
}

func checkNextStatusesIsNotSet(instanceInfo *models.InstanceInfo, statusInfo *models.StatusInfo) (bool, rc.ReturnCode) {

	if statusInfo.Status.StatusType == "OPTIONAL" {
		for _, s := range statusInfo.NextStatuses {
			for _, e := range instanceInfo.Events {
				if e.StatusID == s.ID {
					return false, rc.AT_LEAST_ONE_NEXT_STATUS_IS_SET
				}
			}
		}
	}

	return true, rc.SUCCESS
}

func checkStatusIsSet(instanceInfo *models.InstanceInfo, status *models.Status) rc.ReturnCode {

	for _, e := range instanceInfo.Events {
		if e.StatusID == status.ID {
			return rc.STATUS_IS_SET
		}
	}

	return rc.STATUS_IS_NOT_SET
}

func checkStatusIsReadyToSet(tx *gorm.DB, instanceInfo *models.InstanceInfo, statusInfo *models.StatusInfo, instanceToken string, statusName string) rc.ReturnCode {
	//getting instance info (FOR UPDATE MODE)
	rc5 := instanceInfo.GetInstanceInfo(tx, instanceToken, true)
	if rc5 != rc.SUCCESS {
		return rc5
	}

	finishInstanceIfTimeout(tx, instanceInfo)
	if chk7, rc7 := checkInstanceIsFinished(instanceInfo); chk7 {
		return rc7
	}

	//getting status info
	rc6 := statusInfo.GetStatusInfo(tx, statusName, instanceInfo.Instance.ObjectID)
	if rc6 != rc.SUCCESS {
		return rc6
	}

	//checking status is according to instance
	chk0, rc0 := checkStatusIsBelongsToInstance(instanceInfo, statusInfo)
	if !chk0 {
		return rc0
	}

	//checking previos statuses
	chk1, rc1 := checkPreviosStatusesIsSet(instanceInfo, statusInfo)
	if !chk1 {
		return rc1
	}

	//checking next statuses
	chk2, rc2 := checkNextStatusesIsNotSet(instanceInfo, statusInfo)
	if !chk2 {
		return rc2
	}

	//cheking current status is not set yet
	rc3 := checkStatusIsSet(instanceInfo, &statusInfo.Status)
	if rc3 == rc.STATUS_IS_SET {
		return rc.STATUS_IS_ALREADY_SET
	}

	return rc.SUCCESS
}
