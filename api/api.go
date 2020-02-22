package api

import (
	"encoding/json"
	"net/http"

	"github.com/aleksaan/statusek/logic"
	"github.com/aleksaan/statusek/utils"
)

//---------------------------------------------------------------------------

type ApiCreateInstanceParams struct {
	ObjectName string `json:"object_name"`
}

//ApiCreateInstance - creating instance of specified object type
var ApiCreateInstance = func(w http.ResponseWriter, r *http.Request) {

	params := &ApiCreateInstanceParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	token, err := logic.CreateInstance(params.ObjectName)

	var resp map[string]interface{}
	if err != nil {
		resp = utils.Message(false, err.Error())
	} else {
		resp = utils.Message(true, "success")
	}
	resp["instance_token"] = token
	utils.Respond(w, resp)
}

//---------------------------------------------------------------------------

type ApiSetStatusParams struct {
	InstanceToken string `json:"instance_token"`
	StatusName    string `json:"status_name"`
}

//ApiSetStatus - set status for instance
var ApiSetStatus = func(w http.ResponseWriter, r *http.Request) {

	params := &ApiSetStatusParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	instance, err1 := logic.GetInstanceByToken(params.InstanceToken)
	if err1 != nil {
		utils.Respond(w, utils.Message(false, err1.Error()))
		return
	}

	statusID, err2 := logic.GetStatusIDByName(params.StatusName, instance.ObjectID)
	if err2 != nil {
		utils.Respond(w, utils.Message(false, err2.Error()))
		return
	}

	err3 := logic.SetStatus(instance.InstanceID, statusID)
	if err3 != nil {
		utils.Respond(w, utils.Message(false, err3.Error()))
		return
	}

	var resp map[string]interface{}
	resp = utils.Message(true, "success")
	utils.Respond(w, resp)
}

//---------------------------------------------------------------------------

type ApiCheckInstanceIsFinishedParams struct {
	InstanceToken string `json:"instance_token"`
}

//ApiCheckInstanceIsFinished - checks instanse is finished (return true) or not (return false)
var ApiCheckInstanceIsFinished = func(w http.ResponseWriter, r *http.Request) {

	params := &ApiCheckInstanceIsFinishedParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	instanceID, err1 := logic.GetInstanceIDByToken(params.InstanceToken)
	if err1 != nil {
		utils.Respond(w, utils.Message(false, err1.Error()))
		return
	}

	chk := logic.CheckInstanceIsFinished(instanceID)

	var resp map[string]interface{}
	resp = utils.Message(true, "success")
	resp["is_instance_finished"] = chk
	utils.Respond(w, resp)
}
