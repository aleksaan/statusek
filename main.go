package main

import (
	"fmt"

	"net/http"

	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

var CurrentVersion = "2.0.0"
var db *gorm.DB

func init() {
	db = database.DB
}

func main() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	models.UpdateDB(CurrentVersion)

	// servicePort := os.Getenv("service_port")

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", homeLink)
	// router.HandleFunc("/instance/create", api.ApiCreateInstance)
	// router.HandleFunc("/status/setStatus", api.ApiSetStatus)
	// router.HandleFunc("/instance/checkIsFinished", api.ApiCheckInstanceIsFinished)
	// router.HandleFunc("/instance/getInfo", api.ApiGetInstanceInfo)
	// router.HandleFunc("/event/getEvents", api.ApiGetEvents)
	// router.HandleFunc("/status/checkStatusIsSet", api.ApiCheckStatusIsSet)
	// router.HandleFunc("/status/checkStatusIsReadyToSet", api.ApiCheckStatusIsReadyToSet)
	// router.HandleFunc("/about/", api.ApiAbout)

	// if os.Getenv("ASPNETCORE_PORT") != "" {
	// 	servicePort = os.Getenv("ASPNETCORE_PORT")
	// }

	// http.ListenAndServe(":"+servicePort, router)

}
