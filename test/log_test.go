package test

import (
	"fmt"
	"github.com/DolphinDong/ops-tools/tools"
	"path/filepath"
	"testing"
)

func TestLogger(t *testing.T) {
	logger, err := tools.NewLogger("a/b/c/d/1.log", tools.NO_STDOUT)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	logger.Infof("okok")
}

func TestFilePath(tt *testing.T) {

	fmt.Println(filepath.Dir("a/b/c.log"))
}
