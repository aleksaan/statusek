package main

import (
	"fmt"
	//"log"
	"net/http"

	"github.com/aleksaan/statusek/logic"
	//"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", homeLink)
	//router.HandleFunc("//new", controllers.RestCreateTask)
	// logic.Model()
	logic.SetStatus(1, 1)
	//log.Fatal(http.ListenAndServe(":8080", router))
}
