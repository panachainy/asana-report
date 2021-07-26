package model

import (
	"encoding/json"
	"log"
)

type Tasks struct {
	Data []DataTask `json:"data"`
}

type DataTask struct {
	Gid          string    `json:"gid"`
	Name         string    `json:"name"`
	ResourceType string    `json:"resource_type"`
	Completed    bool      `json:"completed"`
	Assignee     *Assignee `json:"assignee"`
	NumSubTask   int       `json:"num_subtasks"`
}

type Assignee struct {
	Gid          string `json:"gid"`
	ResourceType string `json:"resource_type"`
}

func GetTasksBy(blob string) Tasks {
	var taskResult Tasks

	if err := json.Unmarshal([]byte(blob), &taskResult); err != nil {
		log.Fatal(err)
	}

	return taskResult
}
