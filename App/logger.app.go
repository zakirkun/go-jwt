package app

import (
	"github.com/sirupsen/logrus"
	"os"
)

var L = logrus.New()

func Logger() {
	L.Info("Open logrus")

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	L.SetOutput(file)
	L.Info("Set output to app.log")

	if err != nil {
		L.Info("Failed to log to file, using default stderr")
	}

	L.SetFormatter(&logrus.JSONFormatter{})
	L.Info("Set formatter to json formatter")
}
