package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/aleksaan/scheduler/database"
	"github.com/aleksaan/scheduler/models"
	u "github.com/aleksaan/scheduler/utils"
)

//RestRegisterExecutor -
var RestRegisterExecutor = func(w http.ResponseWriter, r *http.Request) {

	executor := &models.Executor{}

	err := json.NewDecoder(r.Body).Decode(executor)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		return
	}

	ctx := context.Background()
	err = executor.Insert(ctx, database.GetDBX(), boil.Infer())

	var resp map[string]interface{}
	if err != nil {
		resp = u.Message(false, err.Error())
	} else {
		resp = u.Message(true, "success")
	}
	resp["executor"] = executor
	u.Respond(w, resp)
}
