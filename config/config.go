package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type DBConfigType struct {
	DbType                 string
	DbHost                 string
	DbPort                 string
	DbName                 string
	DbUser                 string
	DbPass                 string
	DbSslMode              string
	DbSchema               string
	DbUpdateIfOtherVersion bool
}

type ConfigType struct {
	DBConfig    DBConfigType
	ServicePort string
	GithubLink  string
}

var Config *ConfigType

func init() {
	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}
	Config = New()
}

// New returns a new Config struct
func New() *ConfigType {
	return &ConfigType{
		DBConfig: DBConfigType{
			DbType:                 getEnv("db_type", ""),
			DbHost:                 getEnv("db_host", ""),
			DbPort:                 getEnv("db_port", ""),
			DbName:                 getEnv("db_name", ""),
			DbUser:                 getEnv("db_user", ""),
			DbPass:                 getEnv("db_pass", ""),
			DbSslMode:              getEnv("db_sslmode", ""),
			DbSchema:               getEnv("db_schema", ""),
			DbUpdateIfOtherVersion: getEnvAsBool("db_update_if_other_version", false),
		},
		ServicePort: getEnv("service_port", ""),
		GithubLink:  "https://github.com/aleksaan/statusek",
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
