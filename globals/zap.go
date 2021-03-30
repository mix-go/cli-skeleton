package globals

import (
	"github.com/mix-go/xdi"
	"go.uber.org/zap"
)

func Zap() (logger *zap.SugaredLogger) {
	if err := xdi.Populate("zap", &logger); err != nil {
		panic(err)
	}
	return
}
