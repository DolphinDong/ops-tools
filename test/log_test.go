package test

import (
	"fmt"
	"ops-tools/tools"
	"path/filepath"
	"testing"
)

func TestLogger(t *testing.T) {
	logger, err := tools.NewLogger(tools.NO_LOG_TO_FILE)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	logger.Infof("okok")
}

func TestFilePath(tt *testing.T) {

	fmt.Println(filepath.Dir("a/b/c.log"))
}
