package api

import (
	"encoding/json"
	"fmt"

	"github.com/aleksaan/statusek/database"
	"github.com/aleksaan/statusek/models"
	_ "github.com/lib/pq"
)

type status struct {
	id          uint
	object_id   uint
	status_name string
	status_type string
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
	"MANDATORY": "#606060",
	"OPTIONAL":  "#0066CC",
	"FAILED":    "#FF3333",
}

// var ApiGetGraph = func(w http.ResponseWriter, r *http.Request) {
// 	var resp = &tResponseBody{result: make(map[string]interface{})}
// 	apiCommonStart(r)
// 	_, params := apiCommonDecodeParams(r)
// 	var res = get_graph(params.InstanceToken)
// 	resp.result["res"] = res
// 	sendResponse(w, resp, rc.SUCCESS)
// }

func get_graph(instanceToken string) string {
	db := database.DB

	//select instance
	instance := &models.Instance{}
	instance.GetInstance(db, instanceToken, false)

	//select statuses = nodes
	var istatuses []models.Status
	db.Where("object_id=?", instance.ObjectID).Find(&istatuses)

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
	g.Edges = []edge{}
	iworkflows := []models.Workflow{}
	for _, stat := range istatuses {
		result := db.Where("status_prev_id=?", stat.ID).Find(&iworkflows)
		if result.RowsAffected > 0 {
			for _, wf := range iworkflows {
				g.Edges = append(g.Edges, edge{wf.StatusPrevID, wf.StatusNextID})
			}
		}
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
