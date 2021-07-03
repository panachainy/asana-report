package cmd

import (
	"asana-report/model"
	"asana-report/service"
	"asana-report/util"

	"github.com/spf13/cobra"
)

var isFullReport bool

var astCmd = &cobra.Command{
	Use:   "ast",
	Short: "Get task status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		service.InitService()
		var response model.AstResponse

		workspaceId := util.GLOBAL_CONFIG.WorkspaceId
		token := util.GLOBAL_CONFIG.Token

		util.PrintConfig(cmd)

		cmd.Println("================================================")

		workspace := service.GetWorkspace(workspaceId, token)

		for _, project := range workspace.Project {
			cmd.Printf("Project: %v in progress...\n", project.Name)

			taskCompleted := 0
			tasks := service.GetTasks(project.Gid, token)

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
	} else {
		cmd.Println("==== Short Report ====")
	}
	cmd.Printf("SumTask: %v\n", response.SumTask)
	cmd.Printf("SumCompleted: %v\n", response.SumCompleted)
}
