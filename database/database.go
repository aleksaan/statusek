package database

import (
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

	connectionstring = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Создать строку подключения
}

func ConnectGorm() {
	conn, err := gorm.Open("postgres", connectionstring)

	if err != nil {
		fmt.Print(err)
	}

	DB = conn
}
