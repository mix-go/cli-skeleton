package globals

import (
	"github.com/mix-go/xdi"
	"xorm.io/xorm"
)

func Xorm() (db *xorm.Engine) {
	if err := xdi.Populate("xorm", &db); err != nil {
		panic(err)
	}
	return
}
