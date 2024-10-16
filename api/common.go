package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aleksaan/statusek/logging"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
)

//---------------------------------------------------------------------------

func apiCommonDecodeParams(w http.ResponseWriter, r *http.Request, resp map[string]interface{}) *apiCommonParams {
	logging.RLogger.Info("Parameters parsing...")
	params := &apiCommonParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		message := fmt.Sprintf("Parameters parsing error: %s", err.Error())
		resp["status"] = false
		resp["message"] = message
		logging.RLogger.Error(message)
		utils.Respond(w, resp)
		logging.RLogger.Info("Finished")
		return nil
	}
	return params
}

func apiCommonStart(r *http.Request) {
	logging.CreateRequestLogger(r.RemoteAddr, r.RequestURI)
	logging.RLogger.Info("Started")
}

func apiCommonFinish(w http.ResponseWriter, rcode rc.ReturnCode, resp map[string]interface{}, x map[string]interface{}) {

	var message = replaceLogsTags(rc.ReturnCodes[rcode], x)

	if rcode != rc.SUCCESS {
		resp["status"] = false
		resp["message"] = message
		logging.RLogger.Error(message)
		utils.Respond(w, resp)
		logging.RLogger.Info("Finished")
		return
	}

	resp["status"] = true
	resp["message"] = message
	utils.Respond(w, resp)
	logging.RLogger.Info("Finished")
}
