package api

type apiCommonParams struct {
	ObjectName      string `json:"object_name"`
	InstanceTimeout int    `json:"instance_timeout"`
	StatusName      string `json:"status_name"`
	InstanceToken   string `json:"instance_token"`
}
