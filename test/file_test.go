package test

import (
	"ops-tools/pkg/envconfig"
	"ops-tools/tools"
	"testing"
)

func TestReadConfig(t *testing.T) {

	err := tools.ReadYaml2Options("./config.yaml", &envconfig.Options)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}
