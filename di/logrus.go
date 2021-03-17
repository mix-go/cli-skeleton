package di

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mix-go/cli"
	"github.com/mix-go/di"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	obj := di.Object{
		Name: "logrus",
		New: func() (i interface{}, e error) {
			logger := logrus.New()

			var fileRotate *rotatelogs.RotateLogs
			if err := di.Populate("file-rotatelogs", &fileRotate); err != nil {
				return nil, err
			}
			writer := io.MultiWriter(os.Stdout, fileRotate)
			logger.SetOutput(writer)
			if cli.App().Debug {
				logger.SetLevel(logrus.DebugLevel)
			}

			return logger, nil
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}
