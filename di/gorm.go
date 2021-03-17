package di

import (
	"github.com/mix-go/di"
	"github.com/mix-go/dotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	obj := di.Object{
		Name: "gorm",
		New: func() (i interface{}, e error) {
			return gorm.Open(mysql.Open(dotenv.Getenv("DATABASE_DSN").String()))
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}
