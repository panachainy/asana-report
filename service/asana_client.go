package service

import (
	"asana-report/model"
	"asana-report/util"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func InitService(asanaUrl ...string) {
	url := util.GLOBAL_CONFIG.AsanaUrl

	if len(asanaUrl) > 0 {
		url = asanaUrl[0]
	}

	client.SetHostURL(url)
}

func GetTasks(projectId string, token string) model.Tasks {
	tasks := model.Tasks{}

	response, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetResult(&tasks).
		SetPathParams(map[string]string{
			"project_id": projectId,
		}).
		SetQueryParams(map[string]string{
			"opt_fields": "completed,name,assignee,num_subtasks",
		}).
		Get("projects/{project_id}/tasks")
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at GetTasks()\n", response.StatusCode())
		panic(errorString)
	}

	return tasks
}

func GetSubTasks(taskId string, token string) model.Tasks {
	tasks := model.Tasks{}

	response, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetResult(&tasks).
		SetPathParams(map[string]string{
			"task_gid": taskId,
		}).
		SetQueryParams(map[string]string{
			"opt_fields": "completed,name,assignee",
		}).
		Get("tasks/{task_gid}/subtasks")
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at GetSubTasks()\n", response.StatusCode())
		panic(errorString)
	}

	return tasks
}

func GetWorkspace(workspaceId string, token string) model.Workspace {
	workspace := model.Workspace{}

	response, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetResult(&workspace).
		SetQueryParams(map[string]string{
			"workspace": workspaceId,
		}).
		Get("projects")
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at getWorkspace()\n", response.StatusCode())
		panic(errorString)
	}

	return workspace
}

func UpdateTasks(taskId string, assigneeId string, token string) {
	var asaaRequest model.AsaaRequest
	asaaRequest.Data.Assignee = assigneeId

	response, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetPathParam("task_gid", taskId).
		SetBody(
			asaaRequest,
		).
		Put("tasks/{task_gid}")
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at UpdateTasks()\n", response.StatusCode())
		panic(errorString)
	}
}
