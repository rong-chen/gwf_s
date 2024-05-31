package initRedis

import (
	"github.com/go-redis/redis"
	viper "gwf/core/InitViper"
	"gwf/global"
)

func ConnectRedis(server viper.Server) {
	global.Redis = redis.NewClient(&redis.Options{Addr: server.Host + ":6379"})
	_, err := global.Redis.Ping().Result()
	if err != nil {
		panic("redis connect error:" + err.Error())
	}
}
