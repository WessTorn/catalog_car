package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func InitLogger() {
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logrus.Fatal(err)
	}

	Log.SetLevel(level)
}
