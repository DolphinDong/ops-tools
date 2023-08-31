package main

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"ops-tools/pkg/envconfig"
	"ops-tools/tools"
)

var (
	envConfigCmd *cobra.Command
	configPath   string
)

func main() {
	InitEnvConfigCmd()
	if err := envConfigCmd.Execute(); err != nil {
		errors.WithStack(err)
	}
}
func InitEnvConfigCmd() {
	envConfigCmd = NewEnvConfigCmd()
	envConfigCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./config.yaml", "config path")
}
func NewEnvConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use: "envconfig",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			err = tools.ReadYaml2Options(configPath, &envconfig.Options)
			if err != nil {
				return errors.WithStack(err)
			}
			envconfig.Logger, err = tools.NewLogger(envconfig.Options.LogPath)
			if err != nil {
				return errors.WithStack(err)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			envconfig.Logger.Info("hahahhahah")
			return nil
		},
	}
}
