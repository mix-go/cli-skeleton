module github.com/mix-go/cli-skeleton

go 1.13

replace (
	github.com/mix-go/cli => ../mix/src/cli
	github.com/mix-go/di => ../mix/src/di
	github.com/mix-go/dotenv => ../mix/src/dotenv
)

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/go-redis/redis/v8 v8.0.0-beta.7
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jinzhu/configor v1.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/mix-go/cli v1.1.0
	github.com/mix-go/console v1.0.24 // indirect
	github.com/mix-go/di v1.1.0
	github.com/mix-go/dotenv v1.0.22
	github.com/mix-go/workerpool v1.0.22 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/tebeka/strftime v0.1.5 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/gorm v0.2.34 // indirect
)
