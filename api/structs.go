package api

type tParams struct {
	ObjectName      string `json:"object_name"`
	InstanceTimeout int    `json:"instance_timeout"`
	StatusName      string `json:"status_name"`
	InstanceToken   string `json:"instance_token"`
	StatusMessage   string `json:"status_message"`
}

type tResp struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
