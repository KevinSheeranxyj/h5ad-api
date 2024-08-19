package rdbs

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func TokenGetUser(token string) (user_id string, err error) {
	user_id, err = RDB().Get(ctx, "token:"+token).Result()
	if err != nil {
		if err != redis.Nil {
			logrus.Error(err)
			return
		}
		return "", nil
	}
	return
}

func TokenSet(token string, user_id int) (res string, err error) {
	res, err = RDB().Set(ctx, "token:"+token, user_id, time.Duration(86400)*time.Second).Result()
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(res)
	return
}
