package database

import (
	"github.com/lib/pq"

	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConnectionSettingsType struct {
	UserName  string
	password  string
	DbName    string
	DbHost    string
	DbPort    string
	DbSslMode string
	DbSchema  string
}

var ConnectionSettings = &ConnectionSettingsType{}

//DB - база данных GORM
var DB *gorm.DB

func init() {
	initDbConnectionSettings()
	InitDBConnection()
}

func initDbConnectionSettings() {
	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	ConnectionSettings.UserName = os.Getenv("db_user")
	ConnectionSettings.password = os.Getenv("db_pass")
	ConnectionSettings.DbName = os.Getenv("db_name")
	ConnectionSettings.DbHost = os.Getenv("db_host")
	ConnectionSettings.DbPort = os.Getenv("db_port")
	ConnectionSettings.DbSslMode = os.Getenv("db_sslmode")
	ConnectionSettings.DbSchema = os.Getenv("db_schema")

}

func doConnectionString() string {
	return fmt.Sprintf("host=%s user=%s port=%s dbname=%s sslmode=%s password=%s", ConnectionSettings.DbHost, ConnectionSettings.UserName, ConnectionSettings.DbPort, ConnectionSettings.DbName, ConnectionSettings.DbSslMode, ConnectionSettings.password)
}

func InitDBConnection() rc.ReturnCode {
	conn, err := gorm.Open("postgres", doConnectionString())

	if err != nil {
		WriteDBError(err)
		return rc.DATABASE_ERROR
	}

	DB = conn
	return rc.SUCCESS
}

func WriteDBError(err error) {
	pqErr := err.(*pq.Error)
	fmt.Printf("%s", pqErr.Code.Name())
}
