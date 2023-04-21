package utils

import (
	"encoding/json"
	"net/http"

	rc "github.com/aleksaan/statusek/returncodes"
)

func Respond(w http.ResponseWriter, data map[string]string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/text")
// 	w.Write([]byte("Success"))
// 	return
// }

func ToString(i interface{}) string {
	s, _ := json.MarshalIndent(&i, "", "  ")
	return string(s)
}

// WriteErrorWithRequestToLog - write retrun code and request enviroment to log
func WriteErrorWithRequestToLog(rc rc.ReturnCode, token string) {

}
