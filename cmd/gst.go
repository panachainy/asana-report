package cmd

import (
	"asana-report/model"
	"asana-report/util"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var gstCmd = &cobra.Command{
	Use:   "gst",
	Short: "Get task status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var response model.GstResponse

		workspaceId := util.CONFIG.WorkspaceId
		token := util.CONFIG.Token

		cmd.Println("Configuration")
		cmd.Printf("WorkspaceId: %v\n", workspaceId)
		cmd.Printf("Token: %v\n", token)
		cmd.Println("================================================")

		workspace := getWorkspace(workspaceId, token)

		for _, project := range workspace.Project {
			cmd.Printf("Project: %v in progress...\n", project.Name)

			taskCompleted := 0
			tasks := getTasks(project.Gid, token)

			for _, task := range tasks.Data {
				if task.Completed {
					taskCompleted++
				}
				// cmd.Printf("  Task name: %v is %v\n", task.Name, task.Completed)
			}

			gstData := model.Gst{ProjectName: project.Name, TotalCompleted: taskCompleted, TotalTask: len(tasks.Data)}
			response.Data = append(response.Data, gstData)

			cmd.Println("Done.")
			cmd.Println("================================================")
		}

		response.SumCompleted, response.SumTask = getSumCompletedAndTask(response.Data)

		cmd.Println("All Done.")

		printReport(cmd, response)
	},
}

func init() {
	rootCmd.AddCommand(gstCmd)
}

func getTasks(projectId string, token string) model.Tasks {
	client := resty.New()

	tasks := model.Tasks{}

	response, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetResult(&tasks).
		Get("https://app.asana.com/api/1.1/projects/" + projectId + "/tasks?opt_fields=completed,name")
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at getTasks()\n", response.StatusCode())
		panic(errorString)
	}

	return tasks
}

func getWorkspace(workspaceId string, token string) model.Workspace {
	client := resty.New()

	workspace := model.Workspace{}

	response, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetResult(&workspace).
		Get("https://app.asana.com/api/1.1/projects?limit=10&workspace=" + workspaceId)
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at getWorkspace()\n", response.StatusCode())
		panic(errorString)
	}

	return workspace
}

func getSumCompletedAndTask(gstList []model.Gst) (int, int) {
	sumCompleted := 0
	sumTask := 0

	for _, gst := range gstList {
		sumCompleted = sumCompleted + gst.TotalCompleted
		sumTask = sumTask + gst.TotalTask
	}
	return sumCompleted, sumTask
}

func printReport(cmd *cobra.Command, response model.GstResponse) {
	if true {

		cmd.Println("==== Full Report ====")

		for _, project := range response.Data {
			cmd.Printf("[Project] %v\n", project.ProjectName)
			cmd.Printf("  TotalTask: %v\n", project.TotalTask)
			cmd.Printf("  TotalCompleted: %v\n", project.TotalCompleted)
			cmd.Println("----------------")
		}

		cmd.Printf("SumTask: %v\n", response.SumTask)
		cmd.Printf("SumCompleted: %v\n", response.SumCompleted)
	} else {

		cmd.Println("==== Short Report ====")

		cmd.Printf("SumTask: %v\n", response.SumTask)
		cmd.Printf("SumCompleted: %v\n", response.SumCompleted)
	}
}
