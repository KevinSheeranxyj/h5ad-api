package rdbs

import (
	"context"
	"fmt"
	"uy0/h5ad/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

var rdb *redis.Client
var err error

// RDB Redis
func RDB() *redis.Client {
	if rdb == nil {
		addr := config.Config.Redis.Hostname + config.Config.Redis.HostPort
		logrus.Info("redis info:", addr)
		fmt.Println("-------------------------------------")
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.Config.Redis.Password,
			DB:       config.Config.Redis.Database,
		})
		Res, err := rdb.Ping(ctx).Result()
		if err != nil || Res != "PONG" {
			panic(err)
		}
	}
	return rdb
}
