package excel_tools

import "github.com/sirupsen/logrus"

var (
	Logger  *logrus.Logger
	Options *Config
)

type Config struct {
	LogPath        string          `yaml:"log_path"`
	InputExcelPath string          `yaml:"input_excel_path"`
	SheetSort      SheetSortConfig `yaml:"sheet_sort_config"`
}

type SheetSortConfig struct {
	NeedCreateSubDir bool              `yaml:"need_create_sub_dir"`
	OutputPath       string            `yaml:"output_path"`
	SortField        string            `yaml:"sort_field"`
	SheetName        string            `yaml:"sheet_name"`
	ExcludeRule      map[string]string `yaml:"exclude_rule"`
}
