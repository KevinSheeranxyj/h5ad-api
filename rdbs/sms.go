package rdbs

import (
	"uy0/h5ad/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func SmsSend(mobile string) (err error) {
	_, err = RDB().LPush(ctx, config.ListSms, mobile).Result()
	if err != nil {
		logrus.Error(err)
	}
	return
}

func SmsGetCode(mobile string) (code string, err error) {
	code, err = RDB().Get(ctx, "code:"+mobile).Result()
	if err != nil {
		if err != redis.Nil {
			logrus.Error(err)
			return
		}
		return "", nil
	}
	return
}
