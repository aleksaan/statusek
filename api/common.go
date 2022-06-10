package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aleksaan/statusek/logging"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/aleksaan/statusek/utils"
)

//---------------------------------------------------------------------------

func apiCommonDecodeParams(r *http.Request) (rc.ReturnCode, *apiCommonParams) {
	logging.RLogger.Info("Parameters parsing...")
	params := &apiCommonParams{}

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

func sendResponse(w http.ResponseWriter, resp *tResponseBody, rcode rc.ReturnCode) {
	resp.message = replaceLogsTags(resp, rc.ReturnCodes[rcode])
	if rcode == rc.SUCCESS {
		resp.status = true
	} else {
		resp.status = false
		logging.RLogger.Error(resp.message)
	}
	x := createMapResult(resp)
	utils.Respond(w, x)
}

func createMapResult(resp *tResponseBody) map[string]interface{} {
	out := make(map[string]interface{})
	out["message"] = resp.message
	out["status"] = resp.status
	out["result"] = resp.result
	out["params"] = resp.params
	return out
}

func replaceLogsTags(resp *tResponseBody, str string) string {
	var res string = str
	for k, v := range resp.params {
		res = strings.ReplaceAll(res, "<"+k+">", fmt.Sprintf("%v", v))
	}
	return res
}

func apiCommonStart(r *http.Request) {
	logging.CreateRequestLogger(r.RemoteAddr, r.RequestURI)
	logging.RLogger.Info("Started")
}
