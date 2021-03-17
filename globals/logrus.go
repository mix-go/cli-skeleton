package globals

import (
	"github.com/mix-go/di"
	"github.com/sirupsen/logrus"
)

func Logrus() *logrus.Logger {
	var logger *logrus.Logger
	if err := di.Populate("logrus", &logger); err != nil {
		panic(err)
	}
	return logger
}
