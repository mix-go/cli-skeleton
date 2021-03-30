package di

import (
	"github.com/mix-go/xdi"
	"github.com/mix-go/dotenv"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func init() {
	obj := xdi.Object{
		Name: "xorm",
		New: func() (i interface{}, e error) {
			return xorm.NewEngine("mysql", dotenv.Getenv("DATABASE_DSN").String())
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}
