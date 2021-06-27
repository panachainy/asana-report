package cmd

import (
	"asana-report/util"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "asar",
	Short: "for generate report of asana application",
	Long: `Can generate report of asana application with cli.
	So you can get all task of all your project by workspace,
	You can get summary of task and status.`,
}

var (
	VERSION string = "development"
	CONFIG  util.Config
)

func Execute(version string) {
	VERSION = version

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.asar.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".asar" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".asar")
	}

	viper.SetEnvPrefix(rootCmd.Use)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	util.BindEnvs(CONFIG)

	err := viper.Unmarshal(&CONFIG)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}
