package main

import (
	"github.com/DolphinDong/ops-tools/pkg/excel_tools"

	"github.com/DolphinDong/ops-tools/tools"
	"github.com/spf13/cobra"
)

var SheetSortCmd *cobra.Command

func InitInitEnvCmd() {
	SheetSortCmd = NewSheetSortCmd()
	ExcelToolsCmd.AddCommand(SheetSortCmd)
}
func NewSheetSortCmd() *cobra.Command {
	return &cobra.Command{
		Use: "sheet-sort",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			err = tools.ReadYaml2Options(ConfigPath, &excel_tools.Options)
			if err != nil {
				return err
			}
			excel_tools.Logger, err = tools.NewLogger(excel_tools.Options.LogPath, tools.NEED_STDOUT)
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return excel_tools.SheetSort()
		},
	}
}
