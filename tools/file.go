package tools

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadYaml2Options(configPath string, opts interface{}) error {
	fileByte, err := os.ReadFile(configPath)
	if err != nil {
		return errors.WithStack(err)
	}
	err = yaml.Unmarshal(fileByte, opts)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
