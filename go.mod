module github.com/mix-go/cli-skeleton

go 1.14

replace (
	github.com/mix-go/xcli => ../mix/src/xcli
	github.com/mix-go/xdi => ../mix/src/xdi
	github.com/mix-go/dotenv => ../mix/src/dotenv
)

require (
	github.com/go-redis/redis/v8 v8.7.1
	github.com/jinzhu/configor v1.2.1
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mix-go/xcli v1.1.2
	github.com/mix-go/xdi v1.1.1
	github.com/mix-go/dotenv v1.0.22
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
)
