package test

import (
	"github.com/DolphinDong/ops-tools/pkg/envconfig"
	"github.com/DolphinDong/ops-tools/tools"
	"testing"
)

func TestReadConfig(t *testing.T) {

	err := tools.ReadYaml2Options("./config.yaml", &envconfig.Options)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}
