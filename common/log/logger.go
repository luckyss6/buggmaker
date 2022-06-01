package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func InitLog() (err error) {
	var (
		logfile *os.File
	)

	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if logfile, err = os.OpenFile("./app.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644); err != nil {
		return
	}
	Logger.SetOutput(logfile)
	return
}
