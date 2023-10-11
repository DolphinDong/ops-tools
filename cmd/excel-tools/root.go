package main

import (
	"github.com/DolphinDong/ops-tools/pkg/excel_tools"
	"github.com/spf13/cobra"
	"log"
)

var (
	ExcelToolsCmd *cobra.Command
	ConfigPath    string
)

func main() {
	InitExcelToolsCmd()
	if err := ExcelToolsCmd.Execute(); err != nil {
		if excel_tools.Logger == nil {
			log.Fatalf("%+v", err)
		} else {
			excel_tools.Logger.Fatalf("%+v", err)
		}
	}

}

func InitExcelToolsCmd() {
	ExcelToolsCmd = NewExcelTools()
	ExcelToolsCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "./config.yaml", "config path")
	InitInitEnvCmd()
}
func NewExcelTools() *cobra.Command {
	return &cobra.Command{
		Use: "excel-tools",
	}
}
