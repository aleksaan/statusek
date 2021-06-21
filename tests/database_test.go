// +build integration
package tests

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/stretchr/testify/assert"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

//TestInitDBConnection - testing connection to DB
func TestInitDBConnection(t *testing.T) {
	//test bad connection get error
	un := config.Config.DBConfig.DbUser
	config.Config.DBConfig.DbUser = "a"
	rc1 := database.InitDBConnection()
	assert.Equal(t, rc1, rc.DATABASE_ERROR, "Wrong connection doesn't return error status")

	//test current connection is fine
	config.Config.DBConfig.DbUser = un
	rc0 := database.InitDBConnection()
	assert.Equal(t, rc0, rc.SUCCESS, "Right connection doesn't return success status")
}

func TestUpdateDbIfDisable(t *testing.T) {
	var db = database.DB
	var version = &models.Version{}
	db.First(&version)
	config.Config.DBConfig.DbUpdateIfOtherVersion = false
	rc1 := models.UpdateDB("999")
	assert.Equal(t, rc1, rc.DB_IS_NOT_UPDATED, "Waited status DB_IS_NOT_UPDATED is not returned")
}
