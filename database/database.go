package database

import (
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/logging"
	rc "github.com/aleksaan/statusek/returncodes"
	"gorm.io/driver/postgres"

	"fmt"
)

// DB - база данных GORM
var DB *gorm.DB
var connectionString string

func init() {
	InitDBConnection()
}

func createConnectionString() {
	path, _ := os.Getwd()
	logging.Info("path=%s", path)
	var c = config.Config
	logging.Info("Connection string: host=%s user=%s port=%s dbname=%s sslmode=%s password=%s", c.DBConfig.DbHost, c.DBConfig.DbUser, c.DBConfig.DbPort, c.DBConfig.DbName, c.DBConfig.DbSslMode, "******")
	connectionString = fmt.Sprintf("host=%s user=%s port=%s dbname=%s sslmode=%s password=%s", c.DBConfig.DbHost, c.DBConfig.DbUser, c.DBConfig.DbPort, c.DBConfig.DbName, c.DBConfig.DbSslMode, c.DBConfig.DbPass)
}

func InitDBConnection() rc.ReturnCode {
	createConnectionString()

	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		//SkipDefaultTransaction: true,
	})

	if err != nil {
		logging.Error("%s", err.Error())
		return rc.DATABASE_ERROR
	}

	DB = conn
	logging.Info("DB connection established")
	return rc.SUCCESS
}
