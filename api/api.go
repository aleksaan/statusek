package api

import (
	"net/http"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/logic"
	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/fatih/structs"
)

//---------------------------------------------------------------------------

// ApiCreateInstance - rest api handler creates instance of specified object

var ApiCreateInstance = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}
	instance_token, rc2 := logic.CreateInstance(params.ObjectName, params.InstanceTimeout)
	if rc2 != rc.SUCCESS {
		sendResponse(w, params, result, rc2)
		return
	}
	result.Data["instance_token"] = instance_token
	sendResponse(w, params, result, rc2)
}

//---------------------------------------------------------------------------

// ApiSetStatus - rest api handler sets status for the instance
var ApiSetStatus = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}

	rc2 := logic.SetStatus(params.InstanceToken, params.StatusName)
	sendResponse(w, params, result, rc2)
}

// ApiCheckInstanceIsFinished - rest api handler checks instance is finished (return true) or not (return false)
var ApiCheckInstanceIsFinished = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}

	_, rc2, ii := logic.GetInstanceInfo(params.InstanceToken)
	if rc2 == rc.SUCCESS {
		result.Data["instanse_is_finished_description"] = ii.Instance.InstanceIsFinishedDescription
		result.Data["instance_is_finished"] = ii.Instance.InstanceIsFinished
	}
	sendResponse(w, params, result, rc2)
}

// ApiGetInstanceInfo - rest api handler return info about process
var ApiGetInstanceInfo = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}
	_, rc2, instanceInfo := logic.GetInstanceInfo(params.InstanceToken)

	if rc2 == rc.SUCCESS {
		result.Data["instance_info"] = &instanceInfo
	}
	sendResponse(w, params, result, rc2)
}

// ApiGetEvents - gets events of instance by it token
var ApiGetEvents = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}

	events, rc2 := logic.GetEvents(params.InstanceToken)
	if rc2 == rc.SUCCESS {
		result.Data["events"] = events
	}

	sendResponse(w, params, result, rc2)
}

// ApiCheckStatusIsSet - gets events of instance by it token
var ApiCheckStatusIsSet = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}

	_, rc2 := logic.CheckStatusIsSet(params.InstanceToken, params.StatusName)
	if rc2 == rc.STATUS_IS_SET {
		rc2 = rc.SUCCESS
		result.Data["status_is_set"] = true
	}
	if rc2 == rc.STATUS_IS_NOT_SET {
		rc2 = rc.SUCCESS
		result.Data["status_is_set"] = false
	}

	sendResponse(w, params, result, rc2)
}

// ApiCheckStatusIsReadyToSet - rest api handler sets status for the instance
var ApiCheckStatusIsReadyToSet = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}

	rcode := logic.CheckStatusIsReadyToSet(params.InstanceToken, params.StatusName)

	if rcode == rc.SUCCESS {
		result.Data["status_is_ready_to_set"] = true
		result.Data["status_is_ready_to_set_description"] = ""
		sendResponse(w, params, result, rcode)
		return
	}

	if rcode == rc.STATUS_IS_ALREADY_SET || rcode == rc.NOT_ALL_MANDATORY_ARE_SET || rcode == rc.NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET || rcode == rc.NO_ONE_PREVIOS_OPTIONAL_STATUSES_ARE_SET {
		result.Data["status_is_ready_to_set"] = false
		mapParams := structs.Map(params)
		result.Data["status_is_ready_to_set_description"] = rightMessage(rc.ReturnCodes[rcode], mapParams)

		sendResponse(w, params, result, rc.SUCCESS)
		return
	}

	sendResponse(w, params, result, rcode)
}

// ApiAbout - gets info about program
var ApiAbout = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	params := &tParams{}

	result.Data["version"] = models.CurrentVersion
	result.Data["home page"] = config.Config.GithubLink
	sendResponse(w, params, result, rc.SUCCESS)
}

// ApiSetStatus - rest api handler sets status for the instance
var ApiSetGlobalStatus = func(w http.ResponseWriter, r *http.Request) {
	var result = &tResp{Data: make(map[string]interface{})}
	apiCommonStart(r)
	rc1, params := decodeParams(r)
	if rc1 != rc.SUCCESS {
		sendResponse(w, params, result, rc1)
		return
	}

	rc2 := logic.SetGlobalStatus(params.StatusName)
	sendResponse(w, params, result, rc2)
}
