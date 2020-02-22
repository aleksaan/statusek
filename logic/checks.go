package logic

import (
	"fmt"
)

func CheckStatusIsBelongsToInstance() bool {
	return instanceInfo.Instance.ObjectID == statusInfo.Status.ObjectID
}

func CheckPreviosStatusesIsSet() bool {
	fmt.Printf("\n-----------Check previos statuses is set-------------")
	var countPrevMandatory int
	var countPrevMandatoryIsSet int
	var countPrevOptional int
	var countPrevOptionalIsSet int
	for _, s := range statusInfo.PrevStatuses {
		if s.StatusIsMandatory {
			countPrevMandatory++
			for _, e := range instanceInfo.Events {
				if e.StatusID == s.StatusID {
					countPrevMandatoryIsSet++
					break
				}
			}
		} else {
			countPrevOptional++
			for _, e := range instanceInfo.Events {
				if e.StatusID == s.StatusID {
					countPrevOptionalIsSet++
					break
				}
			}
		}
	}

	fmt.Printf("\ncountPrevMandatory: %d", countPrevMandatory)
	fmt.Printf("\ncountPrevMandatoryIsSet: %d", countPrevMandatoryIsSet)
	fmt.Printf("\ncountPrevOptional: %d", countPrevOptional)
	fmt.Printf("\ncountPrevOptionalIsSet: %d\n", countPrevOptionalIsSet)

	if countPrevMandatory > countPrevMandatoryIsSet {
		fmt.Printf("\n---Failed---")
		return false
	}

	if (countPrevMandatory == 0) && (countPrevOptional > 0) && (countPrevOptionalIsSet == 0) {
		fmt.Printf("\n---Failed---")
		return false
	}

	fmt.Printf("\n---Success---")
	return true
}

func CheckNextStatusesIsNotSet() bool {
	fmt.Printf("\n-----------Check next statuses is not set-------------")
	var isNextIsSet bool

	if statusInfo.Status.StatusIsMandatory {
		fmt.Printf("\n---Success---")
		return true
	}

	for _, s := range statusInfo.NextStatuses {
		for _, e := range instanceInfo.Events {
			if e.StatusID == s.StatusID {
				isNextIsSet = true
				break
			}
		}
	}

	if isNextIsSet {
		fmt.Printf("\n---Failed---")
		return false
	}

	fmt.Printf("\n---Success---")
	return true
}

func CheckCurrentStatusIsNotSet() bool {
	fmt.Printf("\n-----------Check current statuses is not set-------------")
	var isCurrentIsSet bool
	for _, e := range instanceInfo.Events {
		if e.StatusID == statusInfo.Status.StatusID {
			isCurrentIsSet = true
			break
		}
	}
	if isCurrentIsSet {
		fmt.Printf("\n---Failed---")
		return false
	}

	fmt.Printf("\n---Success---")
	return true
}
