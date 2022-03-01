package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

type log struct {
	logs *logrus.Logger
}

// 初始化日志
func NewLog() (*logrus.Logger, error) {

	logs := logrus.New()
	logs.Out = os.Stdout
	logs.Formatter = &logrus.JSONFormatter{}
	return logs, nil
}
