package server

import "uy0/h5ad/rdbs"

// RedisLock 上锁
func RedisLock(key string, val string, expire int) (res bool, err error) {
	res, err = rdbs.RedisLock(key, val, expire)
	return
}
