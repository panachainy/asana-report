package cmd

import (
	"asana-report/model"
	"asana-report/util"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var client = resty.New()

var isFullReport bool

var astCmd = &cobra.Command{
	Use:   "ast",
	Short: "Get task status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var response model.AstResponse

		client.SetHostURL(util.GLOBAL_CONFIG.AsanaUrl)

		workspaceId := util.GLOBAL_CONFIG.WorkspaceId
		token := util.GLOBAL_CONFIG.Token

		util.PrintConfig(cmd)

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

			astData := model.Ast{ProjectName: project.Name, TotalCompleted: taskCompleted, TotalTask: len(tasks.Data)}
			response.Data = append(response.Data, astData)

			cmd.Println("Done.")
			cmd.Println("================================================")
		}

		response.SumCompleted, response.SumTask = getSumCompletedAndTask(response.Data)

		cmd.Println("All Done.")

		printReport(cmd, response)
	},
}

func init() {
	rootCmd.AddCommand(astCmd)
	astCmd.Flags().BoolVarP(&isFullReport, "full-report", "f", false, "add -f tag for print full report (default is short report)")
}

func getTasks(projectId string, token string) model.Tasks {
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

func getWorkspace(workspaceId string, token string) model.Workspace {
	workspace := model.Workspace{}

	response, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetResult(&workspace).
		Get("projects?limit=10&workspace=" + workspaceId)
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != http.StatusOK {
		errorString := fmt.Sprintf("Something wrong from asana status code is %v at getWorkspace()\n", response.StatusCode())
		panic(errorString)
	}

	return workspace
}

func getSumCompletedAndTask(astList []model.Ast) (int, int) {
	sumCompleted := 0
	sumTask := 0

	for _, ast := range astList {
		sumCompleted = sumCompleted + ast.TotalCompleted
		sumTask = sumTask + ast.TotalTask
	}
	return sumCompleted, sumTask
}

func printReport(cmd *cobra.Command, response model.AstResponse) {
	if isFullReport {
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
