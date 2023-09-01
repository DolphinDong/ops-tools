package tools

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	_ "github.com/zput/zxcTool/ztLog/zt_formatter"
	"io"
	"os"
	"path/filepath"
)

const (
	NO_LOG_TO_FILE = ""
	NEED_STDOUT    = true
	NO_STDOUT      = false
)

func NewLogger(logPath string, needStdout bool) (*logrus.Logger, error) {
	if logPath == "" && !needStdout {
		return nil, errors.New("logger output not set")
	}
	logger := logrus.New()
	logger.Formatter = &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
		NoFieldsColors:  true,
		ShowFullLevel:   true,
	}

	if logPath != "" {
		// 创建日志存放的目录
		_ = os.MkdirAll(filepath.Dir(logPath), 0775)
		// 设置日志打印位置
		logfile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if needStdout {
			logger.SetOutput(io.MultiWriter(logfile, os.Stdout))
		} else {
			logger.SetOutput(io.MultiWriter(logfile))
		}
	} else {
		if needStdout {
			logger.SetOutput(os.Stdout)
		}
	}
	return logger, nil
}
