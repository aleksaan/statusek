package main

import (
	"fmt"
	"log"

	//"log"
	"net/http"

	"github.com/aleksaan/statusek/api"
	"github.com/gorilla/mux"
	//"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/instance/create", api.ApiCreateInstance)
	router.HandleFunc("/status/setStatus", api.ApiSetStatus)
	router.HandleFunc("/instance/checkIsFinished", api.ApiCheckInstanceIsFinished)

	log.Fatal(http.ListenAndServe(":8080", router))
}
