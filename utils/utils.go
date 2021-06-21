package utils

import (
	"encoding/json"
	"net/http"

	rc "github.com/aleksaan/statusek/returncodes"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ToString(i interface{}) string {
	s, _ := json.MarshalIndent(&i, "", "  ")
	return string(s)
}

//WriteErrorWithRequestToLog
//write retrun code and request enviroment to log
func WriteErrorWithRequestToLog(rc rc.ReturnCode, token string) {

}
