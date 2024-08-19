package router

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func plaintextToByte(plaintextToByte []byte) (param map[string]interface{}, err error) {
	if err != nil {
		logrus.Error(err)
		return
	}
	param = make(map[string]interface{})
	if err = json.Unmarshal(plaintextToByte, &param); err != nil { // 反序列化为map
		logrus.Error(err)
		return
	}
	return
}
