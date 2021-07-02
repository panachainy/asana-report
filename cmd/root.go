package cmd

import (
	"asana-report/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	VERSION string = "development"
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "asar",
		Short: "for generate report of asana application",
		Long: `Can generate report of asana application with cli.
	So you can get all task of all your project by workspace,
	You can get summary of task and status.`,
	}
)

func Execute(version string) {
	VERSION = version

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.asar.yaml)")
	rootCmd.PersistentFlags().BoolP("full-report", "f", false, "add -f tag for print full report (default is short report)")

	viper.BindPFlag("is_full_report", rootCmd.PersistentFlags().Lookup("full-report"))
	viper.SetDefault("is_full_report", false)
}

func initConfig() {
	util.Init(cfgFile, rootCmd.Use)
}
