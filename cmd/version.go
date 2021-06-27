package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = createVersionCmd()

func createVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "version of your asana-report cli",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(rootCmd.Use + " " + VERSION)
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
