package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var asaaCmd = &cobra.Command{
	Use:   "asaa",
	Short: "Assign all task in assignee with your assigneeId.",
	Long:  `Assign all task in assignee with your assigneeId.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("asao called")
	},
}

func init() {
	rootCmd.AddCommand(asaaCmd)
}
