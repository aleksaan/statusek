// +build integration
package tests

import (
	"testing"

	"github.com/aleksaan/statusek/database"
	rc "github.com/aleksaan/statusek/returncodes"
	"github.com/stretchr/testify/assert"
)

//TestInitDBConnection - testing connection to DB
func TestInitDBConnection(t *testing.T) {
	//test bad connection get error
	un := database.ConnectionSettings.UserName
	database.ConnectionSettings.UserName = "a"
	rc1 := database.InitDBConnection()
	assert.Equal(t, rc1, rc.DATABASE_ERROR, "Wrong connection doesn't return error status")

	//test current connection is fine
	database.ConnectionSettings.UserName = un
	rc0 := database.InitDBConnection()
	assert.Equal(t, rc0, rc.SUCCESS, "Right connection doesn't return success status")
}
