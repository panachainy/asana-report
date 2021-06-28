package cmd

import (
	"github.com/spf13/cobra"
)

var gstCmd = &cobra.Command{
	Use:   "gst",
	Short: "Get task status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		process(cmd)
	},
}

func init() {
	rootCmd.AddCommand(gstCmd)
}

func process(cmd *cobra.Command) {
}
