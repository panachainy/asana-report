package model

type Tasks struct {
	Data []DataTask `json:"data"`
}

type DataTask struct {
	Gid          string    `json:"gid"`
	Name         string    `json:"name"`
	ResourceType string    `json:"resource_type"`
	Completed    bool      `json:"completed"`
	Assignee     *Assignee `json:"assignee"`
	SumSubTask   int       `json:"num_subtasks"`
}

type Assignee struct {
	Gid          string `json:"gid"`
	ResourceType string `json:"resource_type"`
}
