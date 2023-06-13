package tests

import (
	"testing"

	"github.com/aleksaan/statusek/logic"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/stretchr/testify/assert"
)

// TestCreateInstance - creating process (existing & unexiting processes)
func TestCreateInstance(t *testing.T) {

	//Check 1. Non existing object
	_, rc0 := logic.CreateInstance("xxx", 10)

	if rc0 != rc.OBJECT_NAME_IS_NOT_FOUND {
		t.Errorf("Check creating instance of inexistable object: Waited: %s, Got: %s", rc.ReturnCodes[rc.OBJECT_NAME_IS_NOT_FOUND], rc.ReturnCodes[rc0])
	}

	//Check 2. Existing object
	token, rc1 := logic.CreateInstance("2-POINT LINE TASK", 10)

	if rc1 != rc.SUCCESS {
		t.Errorf("Check creating instance of existable object: Waited: %s, Got: %s", rc.ReturnCodes[rc.SUCCESS], rc.ReturnCodes[rc1])
	}

	//Check 3. Token is empty
	if len(token) == 0 {
		t.Errorf("Check token is empty: Waited: false, Got: true")
	}
}

// TestComplexLogic - set of the tests of logic proccess
func TestComplexLogic(t *testing.T) {
	token, _ := logic.CreateInstance("2-POINT LINE TASK", 1000)

	events, _ := logic.GetEvents(token)
	assert.Equal(t, len(events), 0, "Unexpectable one or more events in new created instance")

	//Service doesn't allow set up inexistable status
	rc1 := logic.SetStatus(token, "WRONG", "")
	assert.Equal(t, rc1, rc.STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT, "Setting up unexistable status & didn't get STATUS_NAME_IS_NOT_FOUND_FOR_OBJECT")

	//Set FINISHED status before STARTED and wait NOT_ALL_PREVIOS_MANDATORY_STATUSES_IS_SET
	rc1 = logic.SetStatus(token, "FINISHED", "")
	assert.Equal(t, rc1, rc.NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET, "Set FINISHED before STARTED & didn't get NOT_ALL_PREVIOS_MANDATORY_STATUSES_IS_SET")

	_, rc1 = logic.CheckStatusIsSet(token, "STARTED")
	assert.Equal(t, rc1, rc.STATUS_IS_NOT_SET, "Check status STARTED is set & didn't get STATUS_IS_NOT_SET")

	//Set right status and wait SUCCESS
	rc1 = logic.SetStatus(token, "STARTED", "")
	assert.Equal(t, rc1, rc.SUCCESS, "Set STARTED status & didn't get SUCCESS")

	_, rc1 = logic.CheckStatusIsSet(token, "STARTED")
	assert.Equal(t, rc1, rc.STATUS_IS_SET, "Check status STARTED is set & didn't get SUCCESS")

	_, rc2, ii := logic.GetInstanceInfo(token)
	assert.Equal(t, rc2, rc.SUCCESS, "Call GetInstanceInfo & didn't return SUCCESS")
	assert.Equal(t, ii.Instance.InstanceIsFinished, false, "Call GetInstanceInfo & didn't get InstanceIsFinished=false")

	//Set right status and wait FINISHED
	rc1 = logic.SetStatus(token, "FINISHED", "")
	assert.Equal(t, rc1, rc.SUCCESS, "Set FINISHED status & didn't get SUCCESS")

	events, _ = logic.GetEvents(token)
	assert.Equal(t, 2, len(events), "Expect strongly 2 events")

}
