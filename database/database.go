package database

import (
	log "github.com/aleksaan/statusek/logging"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//var db *gorm.DB //база данных
var connectionstring string

//DB - база данных GORM
var DB *gorm.DB

func init() {
	initConnectionString()
	ConnectGorm()
}

func initConnectionString() {
	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbSslmode := os.Getenv("db_sslmode")

	connectionstring = fmt.Sprintf("host=%s user=%s port=%s dbname=%s sslmode=%s password=%s", dbHost, username, dbPort, dbName, dbSslmode, password) //Создать строку подключения
}

func ConnectGorm() {
	conn, err := gorm.Open("postgres", connectionstring)

	if err != nil {
		log.Error(log.ConfigError, err)
	}

	DB = conn
}
