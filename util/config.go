package util

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GLOBAL_CONFIG = GlobalConfig{
	AsanaUrl: "https://app.asana.com/api/1.1",
}

type GlobalConfig struct {
	WorkspaceId string `mapstructure:"workspace_id"`
	Token       string `mapstructure:"token"`
	AsanaUrl    string `mapstructure:"asana_url"`
	AssigneeId  string `mapstructure:"assignee_id"`
}

func Init(cfgFile string, prefix string) {
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

	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	bindEnvs(GLOBAL_CONFIG)

	err := viper.Unmarshal(&GLOBAL_CONFIG)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
	}
}

func bindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(v.Interface(), append(parts, tv)...)
		default:
			err := viper.BindEnv(strings.Join(append(parts, tv), "."))
			if err != nil {
				fmt.Printf("can't bind config from ENV, %v\n", err)
			}
		}
	}
}

func PrintConfig(cmd *cobra.Command) {
	cmd.Println("[Configuration]")
	cmd.Printf("  GLOBAL_CONFIG: %+v\n", GLOBAL_CONFIG)
}
