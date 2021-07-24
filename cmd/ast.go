package cmd

import (
	"asana-report/model"
	"asana-report/service"
	"asana-report/util"

	"github.com/spf13/cobra"
)

var isFullReportAST bool

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
			countSubTask := 0
			countSubTaskCompleted := 0
			tasks := service.GetTasks(project.Gid, token)

			for _, task := range tasks.Data {

				if task.Completed {
					taskCompleted++
				}

				if isFullReportAST {
					cmd.Printf("  Task name: %v is %v\n", task.Name, task.Completed)
				}

				if task.NumSubTask != 0 {
					subTasks := service.GetSubTasks(task.Gid, token)

					for _, subTask := range subTasks.Data {
						countSubTask++
						if subTask.Completed {
							countSubTaskCompleted++
						}
					}
				}
			}

			astData := model.Ast{ProjectName: project.Name, TotalCompleted: taskCompleted, TotalTask: len(tasks.Data), TotalSubTask: countSubTask, TotalSubTaskCompleted: countSubTaskCompleted}
			response.Data = append(response.Data, astData)

			cmd.Println("Done.")
			cmd.Println("================================================")
		}

		response.SumCompleted, response.SumTask, response.SumSubTask, response.SumSubTaskCompleted = getSumCompletedAndTask(response.Data)

		cmd.Println("All Done.")

		printReport(cmd, response)
	},
}

func init() {
	rootCmd.AddCommand(astCmd)
	astCmd.Flags().BoolVarP(&isFullReportAST, "full-report", "f", false, "add -f tag for print full report (default is short report)")
}

func getSumCompletedAndTask(astList []model.Ast) (int, int, int, int) {
	sumCompleted := 0
	sumTask := 0
	sumSubTask := 0
	sumSubTaskCompleted := 0

	for _, ast := range astList {
		sumCompleted = sumCompleted + ast.TotalCompleted
		sumTask = sumTask + ast.TotalTask
		sumSubTask = sumSubTask + ast.TotalSubTask
		sumSubTaskCompleted = sumSubTaskCompleted + ast.TotalSubTaskCompleted
	}

	return sumCompleted, sumTask, sumSubTask, sumSubTaskCompleted
}

func printReport(cmd *cobra.Command, response model.AstResponse) {
	if isFullReportAST {
		cmd.Println("==== Full Report ====")

		for _, project := range response.Data {
			cmd.Printf("[Project] %v\n", project.ProjectName)
			cmd.Printf("  TotalTask: %v\n", project.TotalTask)
			cmd.Printf("  TotalCompleted: %v\n", project.TotalCompleted)

			cmd.Println("--")

			cmd.Printf("  TotalSubTask: %v\n", project.TotalSubTask)
			cmd.Printf("  TotalSubTaskCompleted: %v\n", project.TotalSubTaskCompleted)

			cmd.Println("----------------")
		}
	} else {
		cmd.Println("==== Short Report ====")
	}
	cmd.Printf("SumTask: %v\n", response.SumTask)
	cmd.Printf("SumCompleted: %v\n", response.SumCompleted)
	cmd.Printf("SumSubTask: %v\n", response.SumSubTask)
	cmd.Printf("SumSubTaskCompleted: %v\n", response.SumSubTaskCompleted)
}
