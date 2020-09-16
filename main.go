package main

import (
	"fmt"
	"os"

	//"log"
	"net/http"

	"github.com/aleksaan/statusek/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	//"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

var Version string

func main() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	servicePort := os.Getenv("service_port")
	Version = os.Getenv("version")

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

	if os.Getenv("ASPNETCORE_PORT") != "" {
		servicePort = os.Getenv("ASPNETCORE_PORT")
	}

	http.ListenAndServe(":"+servicePort, router)

}
