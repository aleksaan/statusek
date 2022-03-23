package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/aleksaan/statusek/api"
	"github.com/aleksaan/statusek/api_show_graph"
	"github.com/aleksaan/statusek/logging"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func init() {

}

func main() {

	logging.Info("Starting service...")

	servicePort := os.Getenv("service_port")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/instance/create", api.ApiCreateInstance)
	router.HandleFunc("/status/setStatus", api.ApiSetStatus)
	router.HandleFunc("/instance/checkIsFinished", api.ApiCheckInstanceIsFinished)
	router.HandleFunc("/instance/getInfo", api.ApiGetInstanceInfo)
	router.HandleFunc("/event/getEvents", api.ApiGetEvents)
	router.HandleFunc("/status/checkStatusIsSet", api.ApiCheckStatusIsSet)
	router.HandleFunc("/status/checkStatusIsReadyToSet", api.ApiCheckStatusIsReadyToSet)
	router.HandleFunc("/about/", api.ApiAbout)
	router.HandleFunc("/instance/graph", api_show_graph.GetGraph)

	if os.Getenv("ASPNETCORE_PORT") != "" {
		servicePort = os.Getenv("ASPNETCORE_PORT")
	}

	http.ListenAndServe("127.0.0.1:"+servicePort, router)
}
