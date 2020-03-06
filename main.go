package main

import (
	"fmt"
	"log"
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

func main() {
	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	servicePort := os.Getenv("service_port")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/instance/create", api.ApiCreateInstance)
	router.HandleFunc("/status/setStatus", api.ApiSetStatus)
	router.HandleFunc("/instance/checkIsFinished", api.ApiCheckInstanceIsFinished)
	router.HandleFunc("/event/getEvents", api.ApiGetEvents)

	if os.Getenv("ASPNETCORE_PORT") != "" {
		servicePort = os.Getenv("ASPNETCORE_PORT")
	}

	log.Fatal(http.ListenAndServe(":"+servicePort, router))

}
