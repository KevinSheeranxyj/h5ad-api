package rdbs

import (
	"uy0/h5ad/config"
)

// Config 配置项
func Config(app string) (val map[string]string, err error) {
	val, err = RDB().HGetAll(ctx, config.HashSetting+app).Result()
	return
}

// Config 配置项
func AppConfig(key string) (val string, err error) {
	val, err = RDB().HGet(ctx, config.AppSetting, key).Result()
	return
}
