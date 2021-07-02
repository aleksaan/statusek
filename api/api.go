package api

import (
	"encoding/json"
	"net/http"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/logging"
	"github.com/aleksaan/statusek/logic"
	"github.com/aleksaan/statusek/models"
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
	logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)
	params := &apiCreateInstanceParams{}

	logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	logging.Info("[%s] [%s] Instance of '%s' model creating...", r.RemoteAddr, r.RequestURI, params.ObjectName)
	token, rc0 := logic.CreateInstance(params.ObjectName, params.InstanceTimeout)

	var resp map[string]interface{}
	if rc0 != rc.SUCCESS {
		logging.Error("[%s] [%s] Instance creating... Error: '%s'", r.RemoteAddr, r.RequestURI, rc.ReturnCodes[rc0])
		resp = utils.Message(false, rc.ReturnCodes[rc0])
	} else {
		logging.Info("[%s] [%s] Instance creating... Done. Instance '%s' created.", r.RemoteAddr, r.RequestURI, token)
		resp = utils.Message(true, rc.ReturnCodes[rc0])
	}
	resp["instance_token"] = token

	utils.Respond(w, resp)
	logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

//---------------------------------------------------------------------------

type apiSetStatusParams struct {
	InstanceToken string `json:"instance_token"`
	StatusName    string `json:"status_name"`
}

//ApiSetStatus - rest api handler sets status for the instance
var ApiSetStatus = func(w http.ResponseWriter, r *http.Request) {
	logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	params := &apiSetStatusParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	logging.Info("[%s] [%s] Setting status '%s' for instance '%s'...", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken)
	rc3 := logic.SetStatus(params.InstanceToken, params.StatusName)
	if rc3 != rc.SUCCESS {
		logging.Error("[%s] [%s] Setting status '%s' for instance '%s'... Error: '%s'", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken, rc.ReturnCodes[rc3])
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc3]))
		return
	}
	logging.Info("[%s] [%s] Setting status '%s' for instance '%s'... Done", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken)

	var resp map[string]interface{} = utils.Message(true, rc.ReturnCodes[rc3])
	utils.Respond(w, resp)
	logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

//---------------------------------------------------------------------------

type apiCheckInstanceIsFinishedParams struct {
	InstanceToken string `json:"instance_token"`
}

//ApiCheckInstanceIsFinished - rest api handler checks instance is finished (return true) or not (return false)
var ApiCheckInstanceIsFinished = func(w http.ResponseWriter, r *http.Request) {
	logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	params := &apiCheckInstanceIsFinishedParams{}
	var resp map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	logging.Info("[%s] [%s] Checking instance '%s' is finished...", r.RemoteAddr, r.RequestURI, params.InstanceToken)
	chk, rc0, instanceInfo := logic.GetInstanceInfo(params.InstanceToken)
	if !chk {
		logging.Error("[%s] [%s] Checking instance '%s' is finished...Error: '%s'", r.RemoteAddr, r.RequestURI, params.InstanceToken, rc.ReturnCodes[rc0])
		resp = utils.Message(false, rc.ReturnCodes[rc0])
		utils.Respond(w, resp)
		return
	}

	if !instanceInfo.Instance.InstanceIsFinished {
		logging.Info("[%s] [%s] Instance '%s' is not finished", r.RemoteAddr, r.RequestURI, params.InstanceToken)
		resp = utils.Message(false, rc.ReturnCodes[rc.INSTANCE_IS_NOT_FINISHED])
		utils.Respond(w, resp)
		return
	}
	logging.Info("[%s] [%s] Instance '%s' is finished by '%s'", r.RemoteAddr, r.RequestURI, params.InstanceToken, instanceInfo.Instance.InstanceIsFinishedDescription)

	resp = utils.Message(true, rc.ReturnCodes[rc.INSTANCE_IS_FINISHED])
	resp["instance_is_finished_description"] = &instanceInfo.Instance.InstanceIsFinishedDescription
	utils.Respond(w, resp)
	logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

//---------------------------------------------------------------------------
type apiGetInstanceInfoParams struct {
	InstanceToken string `json:"instance_token"`
}

// ApiGetInstanceInfo - rest api handler return info about process
var ApiGetInstanceInfo = func(w http.ResponseWriter, r *http.Request) {
	logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	params := &apiGetInstanceInfoParams{}

	var resp map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	logging.Info("[%s] [%s] Getting info by instance '%s'...", r.RemoteAddr, r.RequestURI, params.InstanceToken)
	res, rc0, instanceInfo := logic.GetInstanceInfo(params.InstanceToken)

	if rc0 != rc.SUCCESS {
		logging.Error("[%s] [%s] Getting info by instance '%s'...Error: '%s'", r.RemoteAddr, r.RequestURI, params.InstanceToken, rc.ReturnCodes[rc0])
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc0]))
		return
	}

	resp = utils.Message(res, rc.ReturnCodes[rc0])
	if rc0 == rc.SUCCESS {
		resp["instance"] = &instanceInfo.Instance
	}

	utils.Respond(w, resp)
	logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

//---------------------------------------------------------------------------

type apiGetStatusesParams struct {
	InstanceToken string `json:"instance_token"`
}

// ApiGetEvents - gets events of instance by it token
var ApiGetEvents = func(w http.ResponseWriter, r *http.Request) {
	logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	params := &apiGetStatusesParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	logging.Info("[%s] [%s] Getting events for instance '%s'...", r.RemoteAddr, r.RequestURI, params.InstanceToken)
	events, rc0 := logic.GetEvents(params.InstanceToken)
	if rc0 != rc.SUCCESS {
		logging.Error("[%s] [%s] Getting events for instance '%s'... Error: '%s'", r.RemoteAddr, r.RequestURI, params.InstanceToken, rc.ReturnCodes[rc0])
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc0]))
		return
	}
	logging.Info("[%s] [%s] Getting events for instance '%s'...Done", r.RemoteAddr, r.RequestURI, params.InstanceToken)

	var resp map[string]interface{} = utils.Message(true, rc.ReturnCodes[rc0])
	resp["events"] = events
	utils.Respond(w, resp)
	logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

//---------------------------------------------------------------------------

type apiCheckStatusIsSetParams struct {
	InstanceToken string `json:"instance_token"`
	StatusName    string `json:"status_name"`
}

// ApiCheckStatusIsSet - gets events of instance by it token
var ApiCheckStatusIsSet = func(w http.ResponseWriter, r *http.Request) {
	//logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	//logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	params := &apiCheckStatusIsSetParams{}
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	//logging.Info("[%s] [%s] Check status '%s' is set for token '%s'...", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken)
	res, rc0 := logic.CheckStatusIsSet(params.InstanceToken, params.StatusName)
	if rc0 != rc.STATUS_IS_SET && rc0 != rc.STATUS_IS_NOT_SET {
		logging.Error("[%s] [%s] Check status '%s' is set for token '%s'...Error: '%s'", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken, rc.ReturnCodes[rc0])
	} //else {
	//logging.Info("[%s] [%s] Check status '%s' is set for token '%s'...Done", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken)
	//}

	utils.Respond(w, utils.Message(res, rc.ReturnCodes[rc0]))
	//logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

//---------------------------------------------------------------------------

// ApiAbout - gets info about program
var ApiAbout = func(w http.ResponseWriter, r *http.Request) {
	logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	var resp map[string]interface{} = utils.Message(true, "Success")
	resp["version"] = models.CurrentVersion
	resp["home page"] = config.Config.GithubLink
	utils.Respond(w, resp)
	logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

//---------------------------------------------------------------------------

type apiCheckIsStatusReadyToSetParams struct {
	InstanceToken string `json:"instance_token"`
	StatusName    string `json:"status_name"`
}

//ApiCheckStatusIsReadyToSet - rest api handler sets status for the instance
var ApiCheckStatusIsReadyToSet = func(w http.ResponseWriter, r *http.Request) {
	//logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	//logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	params := &apiCheckIsStatusReadyToSetParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	//logging.Info("[%s] [%s] Check status '%s' is ready to set for token '%s'...", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken)
	rc3 := logic.CheckStatusIsReadyToSet(params.InstanceToken, params.StatusName)
	if rc3 != rc.SUCCESS {
		logging.Error("[%s] [%s] Check status '%s' is ready to set for token '%s'... Error: '%s'", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken, rc.ReturnCodes[3])
		utils.Respond(w, utils.Message(false, rc.ReturnCodes[rc3]))
		return
	}
	//logging.Info("[%s] [%s] Check status '%s' is ready to set for token '%s'... Done", r.RemoteAddr, r.RequestURI, params.StatusName, params.InstanceToken)

	var resp map[string]interface{} = utils.Message(true, rc.ReturnCodes[rc3])
	utils.Respond(w, resp)
	//logging.Info("[%s] [%s] Finished", r.RemoteAddr, r.RequestURI)
}

// func parametersParsing(w http.ResponseWriter, r *http.Request) rc.ReturnCode {
// 	logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
// 	params := &apiSetStatusParams{}

// 	err := json.NewDecoder(r.Body).Decode(params)
// 	if err != nil {
// 		logging.Info("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
// 		utils.Respond(w, utils.Message(false, err.Error()))
// 		return rc.PARAMETERS_PARSING_ERROR
// 	}
// 	return rc.SUCCESS
// }
