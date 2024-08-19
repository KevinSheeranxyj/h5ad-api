package common

import (
	"net/http"

	"uy0/h5ad/rdbs"
	. "uy0/h5ad/tools/resp"
)

type ConfigParam struct {
	K string `json:"k"`
}

func Config(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	app := query.Get("app")

	if app != "" {
		config, err := rdbs.Config(app)
		if err != nil {
			return
		}
		RespData(w, 200, "请求成功", config)
		return
	}

	ErrorParam(w, nil)
}
