package test

import (
	"testing"

	"github.com/aleksaan/statusek/logic"
	rc "github.com/aleksaan/statusek/returncodes"
)

func TestAbs(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

//TestCreateInstance - creating process (existing & unexiting processes)
func TestCreateInstance(t *testing.T) {

	//Non existing object
	_, rc0 := logic.CreateInstance("xxx", 10)

	if rc0 != rc.OBJECT_NAME_IS_NOT_FOUND {
		t.Errorf("Check1: Waited: %s, got: %s", rc.ReturnCodes[rc.OBJECT_NAME_IS_NOT_FOUND], rc.ReturnCodes[rc0])
	}

	//Existing object
	token, rc1 := logic.CreateInstance("2-POINT LINE TASK", 10)

	if rc1 != rc.SUCCESS {
		t.Errorf("Check2: Waited: %s, got: %s", rc.ReturnCodes[rc.SUCCESS], rc.ReturnCodes[rc1])
	}

	if len(token) == 0 {
		t.Errorf("Token is wrong")
	}
}
