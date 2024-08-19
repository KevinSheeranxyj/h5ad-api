package server

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath("$HOME/go/src/robot/robot_wit/")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Info(err)
	}
}

// 锁测试
func TestRedisLock(t *testing.T) {
	key := "lock"
	val := "lock"
	expire := 100
	res, err := RedisLock(key, val, expire)
	fmt.Println(res, err)
	return
}
