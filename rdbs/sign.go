package rdbs

import (
	"time"
	"uy0/h5ad/tools/times"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func SignIn(uuid string) (err error) {
	_, err = RDB().SetNX(ctx, "sign:"+uuid, 1, time.Duration(times.Countdown())*time.Second).Result()
	if err != nil {
		logrus.Error(err)
	}
	return
}

func SignStatus(uuid string) (res string, err error) {
	res, err = RDB().Get(ctx, "sign:"+uuid).Result()
	if err != nil {
		if err == redis.Nil {
			return "0", nil
		}
		logrus.Error(err)
	}
	return
}
