package logic

import (
	"time"

	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
)

func checkStatusIsBelongsToInstance(instanceInfo *models.InstanceInfo, statusInfo *models.StatusInfo) (bool, rc.ReturnCode) {
	if instanceInfo.Instance.ObjectID == statusInfo.Status.Object.ID {
		return true, rc.STATUS_IS_ACCORDING_TO_INSTANCE
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

	if instanceInfo.Instance.InstanceIsFinished == true {
		return true, rc.INSTANCE_IS_FINISHED
	}

	return false, rc.INSTANCE_IS_NOT_FINISHED
}

func checkInstanceIsNotTimeout(instanceInfo *models.InstanceInfo) (bool, rc.ReturnCode) {
	t1 := time.Now()
	t2 := *&instanceInfo.Instance.CreatedAt
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

	if statusInfo.Status.StatusType == "STOP-STATUS" {
		if countPrevMandatoryIsSet+countPrevOptionalIsSet > 0 {
			return true, rc.AT_LEAST_ONE_OF_PREVIOS_STATUSES_IS_SET_FOR_STOP_STATUS
		} else {
			return false, rc.NO_ONE_PREVIOS_STATUSES_ARE_SET_FOR_STOP_STATUS
		}
	}

	if countPrevMandatory > countPrevMandatoryIsSet {
		return false, rc.NOT_ALL_PREVIOS_MANDATORY_STATUSES_IS_SET
		//"Не все обязательные статусы предыдущего уровня установлены"
	}

	if (countPrevMandatory == 0) && (countPrevOptional > 0) && (countPrevOptionalIsSet == 0) {
		return false, rc.NO_ONE_PREVIOS_OPTIONAL_STATUSES_IS_SET
		//"Не установлен ни один опциональный статус предыдущего уровня"
	}

	return true, rc.ALL_PREVIOS_STATUSES_IS_SET
}

func checkNextStatusesIsNotSet(instanceInfo *models.InstanceInfo, statusInfo *models.StatusInfo) (bool, rc.ReturnCode) {

	if statusInfo.Status.StatusType == "MANDATORY" {
		return true, rc.NEXT_STATUSES_IS_NOT_SET
	}

	for _, s := range statusInfo.NextStatuses {
		for _, e := range instanceInfo.Events {
			if e.StatusID == s.ID {
				return false, rc.AT_LEAST_ONE_NEXT_STATUS_IS_SET
			}
		}
	}

	return true, rc.NEXT_STATUSES_IS_NOT_SET
}

func checkCurrentStatusIsNotSet(instanceInfo *models.InstanceInfo, statusInfo *models.StatusInfo) (bool, rc.ReturnCode) {

	for _, e := range instanceInfo.Events {
		if e.StatusID == statusInfo.Status.ID {
			return false, rc.CURRENT_STATUS_IS_SET
		}
	}

	return true, rc.CURRENT_STATUS_IS_NOT_SET
}
