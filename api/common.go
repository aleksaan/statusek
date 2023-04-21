package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aleksaan/statusek/logging"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/fatih/structs"
)

// ---------------------------------------------------------------------------

func decodeParams(r *http.Request) (rc.ReturnCode, *tParams) {
	logging.RLogger.Info("Parameters parsing...")
	params := &tParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		message := fmt.Sprintf("Parameters parsing error: %s", err.Error())
		logging.RLogger.Error(message)
		logging.RLogger.Info("Finished")
		return rc.PARAMS_PARSING_IS_FAILED, nil
	}
	logging.RLogger.Info("Finished")
	return rc.SUCCESS, params
}

func sendResponse(w http.ResponseWriter, params *tParams, resp *tResp, rcode rc.ReturnCode) {
	mapParams := structs.Map(params)
	if rcode == rc.SUCCESS {
		resp.Status = true
	} else {
		resp.Status = false
		resp.Message = rightMessage(rc.ReturnCodes[rcode], mapParams)
		logging.RLogger.Error(resp.Message)
	}
	respond(w, resp)
}

func rightMessage(str string, mapParams map[string]interface{}) string {
	var res string = str
	for k, v := range mapParams {
		res = strings.ReplaceAll(res, "<"+k+">", fmt.Sprintf("%v", v))
	}
	return res
}

func respond(w http.ResponseWriter, resp *tResp) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func apiCommonStart(r *http.Request) {
	logging.CreateRequestLogger(r.RemoteAddr, r.RequestURI)
	logging.RLogger.Info("Started")
}
