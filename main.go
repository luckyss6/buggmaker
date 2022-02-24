package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {

	// 初始化日志
	log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

}
