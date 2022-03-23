package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/models"
	rc "github.com/aleksaan/statusek/returncodes"
	_ "github.com/lib/pq"
)

var c = config.Config

type apiGetGraphParams struct {
	InstanceToken string `json:"instance_token"`
}

type status struct {
	id          uint
	object_id   uint
	status_name string
	status_type string
}

type workflow struct {
	id             uint
	status_prev_id uint
	status_next_id uint
}

type event struct {
	id        uint
	status_id uint
}

type node struct {
	Id    uint   `json:"id"`
	Label string `json:"label"`
	Color string `json:"color"`
}

type edge struct {
	From uint `json:"from"`
	To   uint `json:"to"`
}

type graph struct {
	Nodes []node `json:"nodes"`
	Edges []edge `json:"edges"`
}

type state struct {
	Counter int   `json:"counter"`
	Graph   graph `json:"graph"`
}

var color_map = map[string]string{
	"DEFAULT":   "#98ff98",
	"MANDATORY": "606060",
	"OPTIONAL":  "0066CC",
	"FAILED":    "FF3333",
}

var ApiGetGraph = func(w http.ResponseWriter, r *http.Request) {

	apiCommonStart(r)
	params := apiCommonDecodeParams(w, r)
	var res = get_graph(params.InstanceToken)
	resp["res"] = res
	apiCommonFinish(w, rc.SUCCESS)
}

func get_graph(instanceToken string) string {
	db := database.DB

	//select instance
	instance := &models.Instance{}
	instance.GetInstance(db, instanceToken, false)

	//select statuses = nodes
	var istatuses []models.Status
	db.Where("object_id=?", instance.ObjectID).Find(&istatuses)

	//select workflows = edges
	var iworkflows []models.Workflow
	db.Where("object_id=?", instance.ObjectID).Find(&iworkflows)

	//events of instance
	var ievents []models.Event
	db.Where("instance_id=?", instance.ID).Find(&ievents)

	g := graph{}
	st := state{}

	//fill statuses
	statuses := []status{}
	for _, row := range istatuses {
		statuses = append(statuses, status{id: row.ID, object_id: row.ObjectID, status_name: row.StatusName, status_type: row.StatusType})
	}

	//fill events
	events := []event{}
	for _, ev := range ievents {
		events = append(events, event{id: ev.ID, status_id: ev.StatusID})
	}

	//fill edges
	for _, wf := range iworkflows {
		g.Edges = append(g.Edges, edge{wf.StatusPrevID, wf.StatusNextID})
	}

	//graph
	st.Counter = len(statuses)
	for i := range statuses {
		s := statuses[i]
		var node_color = color_map["DEFAULT"]
		for _, el := range events {
			if el.status_id == s.id {
				node_color = color_map[s.status_type]
				break
			}
		}
		g.Nodes = append(g.Nodes, node{s.id, s.status_name, node_color})
	}
	st.Graph = g

	// To json
	s, err := json.Marshal(st)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	res := string(s)

	return res
}
