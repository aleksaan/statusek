package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/logic"
	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/fatih/structs"
)

var resp map[string]interface{} = make(map[string]interface{})

var x map[string]interface{}

func replaceLogsTags(str string) string {
	var res string = str
	for k, v := range x {
		res = strings.ReplaceAll(res, "<"+k+">", fmt.Sprintf("%v", v))
	}
	return res
}

//---------------------------------------------------------------------------

// ApiCreateInstance - rest api handler creates instance of specified object

var ApiCreateInstance = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)
	x = structs.Map(params)
	instance_token, rcode := logic.CreateInstance(params.ObjectName, params.InstanceTimeout)
	resp["instance_token"] = instance_token
	apiCommonFinish(w, rcode)
}

//---------------------------------------------------------------------------

//ApiSetStatus - rest api handler sets status for the instance
var ApiSetStatus = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)
	rcode := logic.SetStatus(params.InstanceToken, params.StatusName)
	_, _, ii := logic.GetInstanceInfo(params.InstanceToken)
	x["InstanceIsFinishedDescription"] = ii.Instance.InstanceIsFinishedDescription
	apiCommonFinish(w, rcode)
}

//---------------------------------------------------------------------------

//ApiCheckInstanceIsFinished - rest api handler checks instance is finished (return true) or not (return false)
var ApiCheckInstanceIsFinished = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)
	_, rcode, _ := logic.GetInstanceInfo(params.InstanceToken)

	apiCommonFinish(w, rcode)
}

//---------------------------------------------------------------------------

// ApiGetInstanceInfo - rest api handler return info about process
var ApiGetInstanceInfo = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)
	_, rcode, instanceInfo := logic.GetInstanceInfo(params.InstanceToken)

	if rcode == rc.SUCCESS {
		resp["instance"] = &instanceInfo.Instance
	}
	apiCommonFinish(w, rcode)
}

//---------------------------------------------------------------------------

// ApiGetEvents - gets events of instance by it token
var ApiGetEvents = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)

	events, rcode := logic.GetEvents(params.InstanceToken)
	if rcode == rc.SUCCESS {
		resp["events"] = events
	}

	apiCommonFinish(w, rcode)
}

//---------------------------------------------------------------------------

// ApiCheckStatusIsSet - gets events of instance by it token
var ApiCheckStatusIsSet = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)

	_, rcode := logic.CheckStatusIsSet(params.InstanceToken, params.StatusName)

	apiCommonFinish(w, rcode)
}

//---------------------------------------------------------------------------

// ApiAbout - gets info about program
var ApiAbout = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)

	resp["version"] = models.CurrentVersion
	resp["home page"] = config.Config.GithubLink
	apiCommonFinish(w, rc.SUCCESS)
}

//---------------------------------------------------------------------------

//ApiCheckStatusIsReadyToSet - rest api handler sets status for the instance
var ApiCheckStatusIsReadyToSet = func(w http.ResponseWriter, r *http.Request) {
	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)

	rcode := logic.CheckStatusIsReadyToSet(params.InstanceToken, params.StatusName)
	resp["status"] = false
	resp["message"] = rc.ReturnCodes[rcode]

	apiCommonFinish(w, rcode)
}
