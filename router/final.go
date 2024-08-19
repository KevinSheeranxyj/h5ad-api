package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"uy0/h5ad/tools/encrypt"

	. "uy0/h5ad/tools/resp"

	"github.com/sirupsen/logrus"
)

type Param struct {
	Plaintext string `json:"plaintext"`
}

func Final(controller func(http.ResponseWriter, map[string]interface{})) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logrus.Error(err)
			ErrorParam(w, err)
			return
		}

		var p Param

		if err = json.Unmarshal(body, &p); err != nil { //
			logrus.Error(err)
			ErrorParam(w, err)
			return
		}
		ciphetextBase64 := p.Plaintext
		// ciphetextBase64 := string(body)                                       //
		encrypt.PrivateKey, err = ioutil.ReadFile("./tools/encrypt/private.pem") //
		if err != nil {
			logrus.Error(err)
			return
		}
		plaintextByte, err := encrypt.RsaDecryptToByte(ciphetextBase64) //
		if err != nil {
			logrus.Error(err)
			ErrorParam(w, err)
			return
		}
		param, err := plaintextToByte(plaintextByte)
		if err != nil {
			logrus.Error(err)
			return
		}
		controller(w, param)
	})
}

func Empty(controller func(http.ResponseWriter, map[string]interface{})) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := make(map[string]interface{})
		param["r"] = r
		controller(w, param)
	})
}
