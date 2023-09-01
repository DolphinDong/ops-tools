package test

import (
	"fmt"
	"ops-tools/resources/bindata"
	"testing"
)

func TestBindata(t *testing.T) {
	dir, err := bindata.AssetDir("resources/envconfig")
	if err != nil {
		t.Fatalf("%+v", err)
	}
	fmt.Println(dir)
}
