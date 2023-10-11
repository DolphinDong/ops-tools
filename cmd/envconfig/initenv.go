package main

import (
	"github.com/DolphinDong/ops-tools/pkg/envconfig"
	"github.com/DolphinDong/ops-tools/tools"
	"github.com/spf13/cobra"
)

var initEnvCmd *cobra.Command

func InitInitEnvCmd() {
	initEnvCmd = NewInitEnvCmd()
	envConfigCmd.AddCommand(initEnvCmd)
}
func NewInitEnvCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "init-env",
		Aliases: []string{"init"},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			err = tools.ReadYaml2Options(configPath, &envconfig.Options)
			if err != nil {
				return err
			}
			envconfig.Logger, err = tools.NewLogger(envconfig.Options.LogPath, tools.NEED_STDOUT)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			envconfig.Logger.Info("I'm init env cmd")
			return nil
		},
	}
}
