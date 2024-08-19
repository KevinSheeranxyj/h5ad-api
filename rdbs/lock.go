package rdbs

import "time"

// RedisLock 锁
func RedisLock(key string, val string, expire int) (res bool, err error) {
	res, err = RDB().SetNX(ctx, "lock:"+key, val, time.Duration(expire)*time.Second).Result()
	return
}
