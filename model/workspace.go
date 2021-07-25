package model

import (
	"encoding/json"
	"log"
)

type Workspace struct {
	Project  []Project   `json:"data"`
	NextPage interface{} `json:"next_page"`
}

type Project struct {
	Gid          string `json:"gid"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

func GetWorkspaceBy(blob string) Workspace {
	var taskResult Workspace

	if err := json.Unmarshal([]byte(blob), &taskResult); err != nil {
		log.Fatal(err)
	}

	return taskResult
}
