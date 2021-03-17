package globals

import (
	"github.com/jinzhu/gorm"
	"github.com/mix-go/di"
)

func Gorm() *gorm.DB {
	var db *gorm.DB
	if err := di.Populate("gorm", &db); err != nil {
		panic(err)
	}
	return db
}
