package cmd

import (
	"asana-report/model"
	"asana-report/service"
	"asana-report/util"

	"github.com/spf13/cobra"
)

// TODO: combine business to one with ast (ast is main)

var (
	isFullReportRASA bool
	isRASA_All       bool
)

var rasaCmd = &cobra.Command{
	Use:   "rasa",
	Short: "Remove Assign all task in assignee with your assigneeId.",
	Long:  `Remove Assign all task in assignee with your assigneeId.`,
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

				if task.Assignee != nil {
					if !isRASA_All {
						if task.Completed {
							continue
						}
					}

					service.UpdateTasks(task.Gid, "null", token)
					taskCompleted++
				}

				if task.NumSubTask != 0 {
					subTasks := service.GetSubTasks(task.Gid, token)

					// TODO: split sub-task and task
					for _, subTask := range subTasks.Data {
						if subTask.Assignee != nil {

							if !isRASA_All {
								if subTask.Completed {
									continue
								}
							}

							service.UpdateTasks(subTask.Gid, "null", token)
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

		response.SumCompleted, response.SumTask, response.SumSubTask, response.SumSubTaskCompleted = getSumCompletedAndTask(response.Data)

		cmd.Println("All Done.")

		if isFullReportRASA {
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
	rootCmd.AddCommand(rasaCmd)
	rasaCmd.Flags().BoolVarP(&isFullReportRASA, "full-report", "f", false, "add -f tag for print full report (default is short report)")
	rasaCmd.Flags().BoolVarP(&isRASA_All, "all", "a", false, "add -a tag for remove all assignee (default is only remove task incomplete)")
}
