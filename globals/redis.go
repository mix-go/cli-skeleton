package globals

import (
	"github.com/go-redis/redis/v8"
	"github.com/mix-go/di"
)

func Redis() (client *redis.Client) {
	if err := di.Populate("redis", &client); err != nil {
		panic(err)
	}
	return
}
