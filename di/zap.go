package di

import (
	"fmt"
	"github.com/mix-go/xcli"
	"github.com/mix-go/xdi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

func init() {
	obj := xdi.Object{
		Name: "zap",
		New: func() (i interface{}, e error) {
			filename := fmt.Sprintf("%s/../logs/cli.log", xcli.App().BasePath)
			fileRotate := &lumberjack.Logger{
				Filename:   filename,
				MaxBackups: 7,
			}
			writer := io.MultiWriter(os.Stdout, fileRotate)
			w := zapcore.AddSync(writer)
			core := zapcore.NewCore(
				zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
					TimeKey:       "T",
					LevelKey:      "L",
					NameKey:       "N",
					CallerKey:     "C",
					MessageKey:    "M",
					StacktraceKey: "S",
					LineEnding:    zapcore.DefaultLineEnding,
					EncodeLevel:   zapcore.CapitalLevelEncoder,
					EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
						enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
					},
					EncodeDuration: zapcore.StringDurationEncoder,
					EncodeCaller:   zapcore.ShortCallerEncoder,
				}),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
					w),
				zap.DebugLevel,
			)
			logger := zap.New(core, zap.AddCaller())
			return logger.Sugar(), nil
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}
