package globals

import (
	"github.com/mix-go/di"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	var db *gorm.DB
	if err := di.Populate("gorm", &db); err != nil {
		panic(err)
	}
	return db
}
