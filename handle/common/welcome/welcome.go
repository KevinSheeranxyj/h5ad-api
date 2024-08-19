package welcome

import (
	"net/http"

	. "uy0/h5ad/tools/resp"

	"github.com/sirupsen/logrus"
)

// Welcome 落地页
func Welcome(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"welcome": "welcome to visit api",
	}

	RespData(w, 200, "successful", data)
}

func Hi(w http.ResponseWriter, param map[string]interface{}) {
	logrus.Traceln(param)
	data := map[string]string{
		"welcome": "welcome to visit api",
	}
	RespData(w, 200, "successful", data)
}
