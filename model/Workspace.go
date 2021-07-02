package model

type Workspace struct {
	Project  []Project   `json:"data"`
	NextPage interface{} `json:"next_page"`
}

type Project struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}
