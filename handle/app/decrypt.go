package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"uy0/h5ad/dao"
	"uy0/h5ad/tools/encrypt"

	"github.com/sirupsen/logrus"
)

type Param struct {
	Plaintext string `json:"plaintext"`
}

// 明文转为结构体
func DecryptToDao(r *http.Request, dao dao.Dao) (err error) {
	plaintextByte, err := Decrypt(r)
	if err != nil {
		logrus.Error(err)
		return
	}
	if err = json.Unmarshal(plaintextByte, dao); err != nil { // 反序列化为结构体
		logrus.Error(err)
		return
	}
	return
}

// 明文转为map
func DecryptToMap(r *http.Request) (param map[string]string, err error) {
	plaintextByte, err := Decrypt(r)
	if err != nil {
		logrus.Error(err)
		return
	}
	param = make(map[string]string)
	if err = json.Unmarshal(plaintextByte, &param); err != nil { // 反序列化为map
		logrus.Error(err)
		return
	}
	return
}

// 参数接收 到解析出明文
func Decrypt(r *http.Request) (plaintextByte []byte, err error) {
	body, err := ioutil.ReadAll(r.Body) // 接收前端参数 byte
	if err != nil {
		logrus.Error(err)
		return
	}

	var p Param

	if err = json.Unmarshal(body, &p); err != nil { // 反序列化为结构体
		logrus.Error(err)
		return
	}
	ciphetextBase64 := p.Plaintext

	// ciphetextBase64 := string(body) // 转化为string 这里前端传递的是base64

	encrypt.PrivateKey, err = ioutil.ReadFile("./tools/encrypt/private.pem") // 私钥证书
	if err != nil {
		logrus.Error(err)
		return
	}
	plaintextByte, err = encrypt.RsaDecryptToByte(ciphetextBase64) // 解密
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}
