package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aleksaan/statusek/logic"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
	"github.com/joho/godotenv"
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

//ApiCheckInstanceIsFinished - rest api handler checks instance is finished (return true) or not (return false)
var ApiCheckInstanceIsFinished = func(w http.ResponseWriter, r *http.Request) {

	params := &apiCheckInstanceIsFinishedParams{}

	var resp map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	chk, rc0, instanceInfo := logic.GetInstanceInfo(params.InstanceToken)
	if chk == false {
		resp = utils.Message(false, rc.ReturnCodes[rc0])
		utils.Respond(w, resp)
		return
	}

	if instanceInfo.Instance.InstanceIsFinished == false {
		resp = utils.Message(false, rc.ReturnCodes[rc.INSTANCE_IS_NOT_FINISHED])
		utils.Respond(w, resp)
		return
	}

	resp = utils.Message(true, rc.ReturnCodes[rc.INSTANCE_IS_FINISHED])
	utils.Respond(w, resp)
}

//---------------------------------------------------------------------------
type apiGetInstanceInfoParams struct {
	InstanceToken string `json:"instance_token"`
}

// ApiGetInstanceInfo - rest api handler return info about process
var ApiGetInstanceInfo = func(w http.ResponseWriter, r *http.Request) {

	params := &apiGetInstanceInfoParams{}

	var resp map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	res, rc0, instanceInfo := logic.GetInstanceInfo(params.InstanceToken)
	resp = utils.Message(res, rc.ReturnCodes[rc0])
	if rc0 == rc.SUCCESS {
		resp["instance"] = &instanceInfo.Instance
	}

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

//---------------------------------------------------------------------------

// ApiAbout - gets info about program
var ApiAbout = func(w http.ResponseWriter, r *http.Request) {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}
	version := os.Getenv("version")
	githublink := os.Getenv("githublink")

	var resp map[string]interface{}
	resp = utils.Message(true, "Success")
	resp["version"] = version
	resp["home page"] = githublink
	utils.Respond(w, resp)
}
