package cmd

import (
	"asana-report/model"
	"asana-report/service"
	"asana-report/util"

	"github.com/spf13/cobra"
)

var asaaCmd = &cobra.Command{
	Use:   "asaa",
	Short: "Assign all task in assignee with your assigneeId.",
	Long:  `Assign all task in assignee with your assigneeId.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.InitService()
		var response model.AstResponse

		workspaceId := util.GLOBAL_CONFIG.WorkspaceId
		token := util.GLOBAL_CONFIG.Token
		assigneeId := util.GLOBAL_CONFIG.AssigneeId

		util.PrintConfig(cmd)

		cmd.Println("================================================")

		workspace := service.GetWorkspace(workspaceId, token)

		for _, project := range workspace.Project {
			cmd.Printf("Project: %v in progress...\n", project.Name)

			taskCompleted := 0
			tasks := service.GetTasks(project.Gid, token)

			for _, task := range tasks.Data {

				if task.Assignee == nil {
					service.UpdateTasks(task.Gid, assigneeId, token)
					taskCompleted++
				}

				if task.SumSubTask != 0 {
					subTasks := service.GetSubTasks(task.Gid, token)

					// TODO: split sub-task and task
					for _, subTask := range subTasks.Data {
						if subTask.Assignee == nil {
							service.UpdateTasks(subTask.Gid, assigneeId, token)
							taskCompleted++
						}
					}
				}
			}

			astData := model.Ast{ProjectName: project.Name, TotalCompleted: taskCompleted, TotalTask: len(tasks.Data)}
			response.Data = append(response.Data, astData)

			cmd.Println("Done.")
			cmd.Println("================================================")
		}

		response.SumCompleted, response.SumTask = getSumCompletedAndTask(response.Data)

		cmd.Println("All Done.")

		if isFullReport {
			cmd.Println("==== Full Report ====")

			for _, project := range response.Data {
				cmd.Printf("[Project] %v\n", project.ProjectName)
				cmd.Printf("  TotalTask: %v\n", project.TotalTask)
				cmd.Printf("  TotalUpdatedAssignee: %v\n", project.TotalCompleted)
				cmd.Println("----------------")
			}
		} else {
			cmd.Println("==== Short Report ====")
		}

		cmd.Printf("SumTask: %v\n", response.SumTask)
		cmd.Printf("SumUpdatedAssignee: %v\n", response.SumCompleted)
	},
}

func init() {
	rootCmd.AddCommand(asaaCmd)
}
