package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {

	logger = logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	//todo: read log path from config file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	} else {
		logger.Info("Failed to log to file, using default stderr")
		os.Exit(-1)
	}
}

func GetLogger() *logrus.Logger {
	return logger
}
