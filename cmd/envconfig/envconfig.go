package main

import (
	"github.com/DolphinDong/ops-tools/pkg/envconfig"
	"github.com/spf13/cobra"
	"log"
)

var (
	envConfigCmd *cobra.Command
	configPath   string
)

func main() {
	InitEnvConfigCmd()
	if err := envConfigCmd.Execute(); err != nil {
		if envconfig.Logger == nil {
			log.Fatalf("%+v", err)
		} else {
			envconfig.Logger.Fatalf("%+v", err)
		}
	}
}
func InitEnvConfigCmd() {
	envConfigCmd = NewEnvConfigCmd()
	envConfigCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./config.yaml", "config path")
	InitInitEnvCmd()
}
func NewEnvConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use: "envconfig",
	}
}
