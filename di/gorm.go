package di

import (
	"github.com/jinzhu/gorm"
	"github.com/mix-go/di"
	"github.com/mix-go/dotenv"
)

func init() {
	obj := di.Object{
		Name: "gorm",
		New: func() (i interface{}, e error) {
			return gorm.Open("mysql", dotenv.Getenv("DATABASE_DSN").String())
		},
	}
	if err := di.Provide(&obj); err != nil {
		panic(err)
	}
}
