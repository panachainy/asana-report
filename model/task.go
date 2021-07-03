package model

type Tasks struct {
	Data []DataTask `json:"data"`
}

type DataTask struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
	Completed    bool   `json:"completed"`
}
