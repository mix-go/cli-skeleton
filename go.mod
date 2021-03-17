module github.com/mix-go/cli-skeleton

go 1.14

replace (
	github.com/mix-go/cli => ../mix/src/cli
	github.com/mix-go/di => ../mix/src/di
	github.com/mix-go/dotenv => ../mix/src/dotenv
)

require (
	github.com/go-redis/redis/v8 v8.7.1
	github.com/jinzhu/configor v1.2.1
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/mix-go/cli v0.0.0-00010101000000-000000000000
	github.com/mix-go/di v0.0.0-00010101000000-000000000000
	github.com/mix-go/dotenv v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.3
)
