package api

import (
	"encoding/json"
	"net/http"

	"github.com/aleksaan/statusek/logic"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
)

//---------------------------------------------------------------------------

type apiCreateInstanceParams struct {
	ObjectName      string `json:"object_name"`
	InstanceTimeout int    `json:"instance_timeout"`
}

// ApiCreateInstance - rest api handler creates instance of specified object

var ApiCreateInstance = func(w http.ResponseWriter, r *http.Request) {

	params := &apiCreateInstanceParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	token, rc0 := logic.CreateInstance(params.ObjectName, params.InstanceTimeout)

	var resp map[string]interface{}
	if rc0 != rc.SUCCESS {
		resp = utils.Message(false, rc.ReturnCodes[rc0])
	} else {
		resp = utils.Message(true, rc.ReturnCodes[rc0])
	}
	resp["instance_token"] = token
	utils.Respond(w, resp)
}

//---------------------------------------------------------------------------

type apiSetStatusParams struct {
	InstanceToken string `json:"instance_token"`
	StatusName    string `json:"status_name"`
}

//ApiSetStatus - rest api handler sets status for the instance
var ApiSetStatus = func(w http.ResponseWriter, r *http.Request) {

	params := &apiSetStatusParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	rc3 := logic.SetStatus(params.InstanceToken, params.StatusName)
	if rc3 != rc.SUCCESS {
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc3]))
		return
	}

	var resp map[string]interface{}
	resp = utils.Message(true, rc.ReturnCodes[rc3])
	utils.Respond(w, resp)
}

//---------------------------------------------------------------------------

type apiCheckInstanceIsFinishedParams struct {
	InstanceToken string `json:"instance_token"`
}

// ApiCheckInstanceIsFinished - rest api handler checks instance is finished (return true) or not (return false)
// Finished is if all of mandatory statuses of last level is set or in case no one mandatory
// then at least one of optional statuses is set
var ApiCheckInstanceIsFinished = func(w http.ResponseWriter, r *http.Request) {

	params := &apiCheckInstanceIsFinishedParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	chk, rc0 := logic.CheckInstanceIsFinished(params.InstanceToken)
	if chk == false {
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc0]))
		return
	}

	var resp map[string]interface{}
	resp = utils.Message(true, rc.ReturnCodes[rc0])
	utils.Respond(w, resp)
}

//---------------------------------------------------------------------------

type apiGetStatusesParams struct {
	InstanceToken string `json:"instance_token"`
}

// ApiGetEvents - gets events of instance by it token
var ApiGetEvents = func(w http.ResponseWriter, r *http.Request) {

	params := &apiGetStatusesParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	events, rc0 := logic.GetEvents(params.InstanceToken)
	if rc0 != rc.SUCCESS {
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc0]))
		return
	}

	var resp map[string]interface{}
	resp = utils.Message(true, rc.ReturnCodes[rc0])
	resp["events"] = events
	utils.Respond(w, resp)
}

//---------------------------------------------------------------------------

type apiCheckStatusIsSetParams struct {
	InstanceToken string `json:"instance_token"`
	StatusName    string `json:"status_name"`
}

// ApiCheckStatusIsSet - gets events of instance by it token
var ApiCheckStatusIsSet = func(w http.ResponseWriter, r *http.Request) {

	params := &apiCheckStatusIsSetParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	res, rc0 := logic.CheckStatusIsSet(params.InstanceToken, params.StatusName)
	if !res {
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc0]))
		return
	}

	var resp map[string]interface{}
	resp = utils.Message(true, rc.ReturnCodes[rc0])
	utils.Respond(w, resp)
}
