package service

import (
	"asana-report/model"
	"asana-report/util"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

func InitService() {
	client.SetHostURL(util.GLOBAL_CONFIG.AsanaUrl)
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
			"opt_fields": "completed,name",
		}).
		Get("projects/{project_id}/tasks")
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at getTasks()\n", response.StatusCode())
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
