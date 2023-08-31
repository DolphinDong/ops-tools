package envconfig

import "github.com/sirupsen/logrus"

var (
	Logger  *logrus.Logger
	Options *Config
)

type Config struct {
	LogPath string `yaml:"log_path"`
}
