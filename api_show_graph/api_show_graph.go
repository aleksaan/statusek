package api_show_graph

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aleksaan/statusek/config"
	"github.com/aleksaan/statusek/logging"
	"github.com/aleksaan/statusek/utils"
	_ "github.com/lib/pq"
)

var c = config.Config

var DB_DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBConfig.DbUser, c.DBConfig.DbPass, c.DBConfig.DbHost, c.DBConfig.DbPort, c.DBConfig.DbName)

type apiGetGraphParams struct {
	InstanceToken string `json:"instance_token"`
}

type status struct {
	id          int
	object_id   string
	status_name string
	status_type string
}

type workflow struct {
	id             int
	status_prev_id int
	status_next_id int
}

type event struct {
	id        int
	status_id int
}

type node struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
	Color string `json:"color"`
}

type edge struct {
	From int `json:"from"`
	To   int `json:"to"`
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

var GetGraph = func(w http.ResponseWriter, r *http.Request) {
	logging.Info("[%s] [%s] Started", r.RemoteAddr, r.RequestURI)

	logging.Info("[%s] [%s] Parameters parsing...", r.RemoteAddr, r.RequestURI)
	params := &apiGetGraphParams{}
	var resp map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(params)

	if err != nil {
		logging.Error("[%s] [%s] Parameters parsing error:", r.RemoteAddr, r.RequestURI, err.Error())
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}
	var res = get_graph(params.InstanceToken)
	resp = utils.Message(true, res)

	utils.Respond(w, resp)
}

func get_graph(instance_token string) string {
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		logging.Error("Failed to open a DB connection: ", err)
	}
	defer db.Close()

	var sreq = fmt.Sprintf("select t3.id, t3.object_id, t3.status_name, t3.status_type from statusek.instances t1 inner join statusek.statuses t3 on (t1.instance_token = '%s' and t3.object_id = t1.object_id)", instance_token)
	srows, error := db.Query(sreq)
	if error != nil {
		logging.Error("Failed to get instanse with id=%s", instance_token)
		return ""
	}

	statuses := []status{}
	for srows.Next() {
		s := status{}
		srows.Scan(&s.id, &s.object_id, &s.status_name, &s.status_type)
		statuses = append(statuses, s)
	}

	var wreq = fmt.Sprintf("with st as (select t1.instance_token, t3.status_name, t3.status_type, t3.id as status_id from statusek.instances t1 inner join statusek.statuses t3 on (t1.instance_token = '%s' and t3.object_id = t1.object_id)) select id, status_prev_id, status_next_id from statusek.workflows where status_prev_id in (select status_id from st)", instance_token)

	wrows, error := db.Query(wreq)

	if error != nil {
		logging.Error("Failed to get instanse with id=%s", instance_token)
		return ""
	}
	workflows := []workflow{}
	for wrows.Next() {
		w := workflow{}
		wrows.Scan(&w.id, &w.status_prev_id, &w.status_next_id)
		workflows = append(workflows, w)
	}

	var ereq = fmt.Sprintf("select t1.* from statusek.events t1 inner join statusek.instances t2 on (t2.instance_token='%s' and t2.id = t1.instance_id)", instance_token)
	erows, error := db.Query(ereq)
	if error != nil {
		logging.Error("Failed to get events with id=%s", instance_token)
		return ""
	}

	events := []event{}
	for erows.Next() {
		e := event{}
		erows.Scan(&e.id, &e.status_id)
		events = append(events, e)
	}

	log.Print(len(statuses))
	//To graph
	st := state{}

	g := graph{}
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

	for i := range workflows {
		w := workflows[i]
		g.Edges = append(g.Edges, edge{w.status_prev_id, w.status_next_id})
	}
	st.Graph = g

	log.Print(st)
	// To json
	s, err := json.Marshal(st)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	res := string(s)
	log.Print(res)
	return res
}
