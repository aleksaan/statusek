package api

import (
	"net/http"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/logic"
	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/fatih/structs"
)

type tResponseBody struct {
	params  map[string]interface{}
	result  map[string]interface{}
	message string
	status  bool
}

//---------------------------------------------------------------------------

// ApiCreateInstance - rest api handler creates instance of specified object

var ApiCreateInstance = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	apiCommonStart(r)
	rcode1, params := apiCommonDecodeParams(r)
	resp.params = structs.Map(params)
	if rcode1 != rc.SUCCESS {
		sendResponse(w, resp, rcode1)
		return
	}
	instance_token, rcode2 := logic.CreateInstance(params.ObjectName, params.InstanceTimeout)
	if rcode2 != rc.SUCCESS {
		sendResponse(w, resp, rcode2)
		return
	}
	resp.result["instance_token"] = instance_token
	sendResponse(w, resp, rcode2)
}

//---------------------------------------------------------------------------

//ApiSetStatus - rest api handler sets status for the instance
var ApiSetStatus = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	apiCommonStart(r)
	rcode1, params := apiCommonDecodeParams(r)
	resp.params = structs.Map(params)
	if rcode1 != rc.SUCCESS {
		sendResponse(w, resp, rcode1)
		return
	}
	rcode := logic.SetStatus(params.InstanceToken, params.StatusName)
	if rcode != rc.SUCCESS {
		if rcode == rc.INSTANCE_IS_FINISHED {
			_, _, ii := logic.GetInstanceInfo(params.InstanceToken)
			resp.params["InstanceIsFinishedDescription"] = ii.Instance.InstanceIsFinishedDescription
		}
		sendResponse(w, resp, rcode)
		return
	}
	sendResponse(w, resp, rcode)
}

//---------------------------------------------------------------------------

//ApiCheckInstanceIsFinished - rest api handler checks instance is finished (return true) or not (return false)
var ApiCheckInstanceIsFinished = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	apiCommonStart(r)
	rcode1, params := apiCommonDecodeParams(r)
	resp.params = structs.Map(params)
	if rcode1 != rc.SUCCESS {
		sendResponse(w, resp, rcode1)
		return
	}
	_, rcode, ii := logic.GetInstanceInfo(params.InstanceToken)
	if rcode == rc.SUCCESS {
		resp.result["instanse_is_finished_description"] = ii.Instance.InstanceIsFinishedDescription
		resp.result["instance_is_finished"] = ii.Instance.InstanceIsFinished
	}
	sendResponse(w, resp, rcode)
}

//---------------------------------------------------------------------------

// ApiGetInstanceInfo - rest api handler return info about process
var ApiGetInstanceInfo = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	apiCommonStart(r)
	_, params := apiCommonDecodeParams(r)
	_, rcode, instanceInfo := logic.GetInstanceInfo(params.InstanceToken)

	if rcode == rc.SUCCESS {
		resp.result["instance"] = &instanceInfo.Instance
	}
	sendResponse(w, resp, rcode)
}

//---------------------------------------------------------------------------

// ApiGetEvents - gets events of instance by it token
var ApiGetEvents = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	apiCommonStart(r)
	rcode1, params := apiCommonDecodeParams(r)
	resp.params = structs.Map(params)
	if rcode1 != rc.SUCCESS {
		sendResponse(w, resp, rcode1)
		return
	}

	events, rcode := logic.GetEvents(params.InstanceToken)
	if rcode == rc.SUCCESS {
		resp.result["events"] = events
	}

	sendResponse(w, resp, rcode)
}

//---------------------------------------------------------------------------

// ApiCheckStatusIsSet - gets events of instance by it token
var ApiCheckStatusIsSet = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	apiCommonStart(r)
	rcode1, params := apiCommonDecodeParams(r)
	resp.params = structs.Map(params)
	if rcode1 != rc.SUCCESS {
		sendResponse(w, resp, rcode1)
		return
	}

	_, rcode := logic.CheckStatusIsSet(params.InstanceToken, params.StatusName)
	if rcode == rc.STATUS_IS_SET {
		rcode = rc.SUCCESS
		resp.result["status_is_set"] = true
	}
	if rcode == rc.STATUS_IS_NOT_SET {
		rcode = rc.SUCCESS
		resp.result["status_is_set"] = false
	}
	if rcode != rc.SUCCESS {
		if rcode == rc.INSTANCE_IS_FINISHED {
			_, _, ii := logic.GetInstanceInfo(params.InstanceToken)
			resp.params["InstanceIsFinishedDescription"] = ii.Instance.InstanceIsFinishedDescription
		}
		sendResponse(w, resp, rcode)
		return
	}
	sendResponse(w, resp, rcode)
}

//---------------------------------------------------------------------------

// ApiAbout - gets info about program
var ApiAbout = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	resp.params = make(map[string]interface{})
	resp.result["result"] = nil
	apiCommonStart(r)
	resp.result["version"] = models.CurrentVersion
	resp.result["home page"] = config.Config.GithubLink
	sendResponse(w, resp, rc.SUCCESS)
}

//---------------------------------------------------------------------------

//ApiCheckStatusIsReadyToSet - rest api handler sets status for the instance
var ApiCheckStatusIsReadyToSet = func(w http.ResponseWriter, r *http.Request) {
	var resp = &tResponseBody{result: make(map[string]interface{})}
	apiCommonStart(r)
	rcode1, params := apiCommonDecodeParams(r)
	resp.params = structs.Map(params)
	if rcode1 != rc.SUCCESS {
		sendResponse(w, resp, rcode1)
		return
	}

	rcode := logic.CheckStatusIsReadyToSet(params.InstanceToken, params.StatusName)

	if rcode == rc.SUCCESS {
		resp.result["status_is_ready_to_set"] = true
		resp.result["status_is_ready_to_set_description"] = ""
		sendResponse(w, resp, rcode)
		return
	}

	if rcode == rc.NOT_ALL_MANDATORY_ARE_SET || rcode == rc.NOT_ALL_PREVIOS_MANDATORY_STATUSES_ARE_SET || rcode == rc.NO_ONE_PREVIOS_OPTIONAL_STATUSES_ARE_SET {
		resp.result["status_is_ready_to_set"] = false
		resp.result["status_is_ready_to_set_description"] = replaceLogsTags(resp, rc.ReturnCodes[rcode])
		sendResponse(w, resp, rc.SUCCESS)
		return
	}

	sendResponse(w, resp, rcode)
}
