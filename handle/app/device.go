package app

import (
	"net/http"

	"uy0/h5ad/dao"
	. "uy0/h5ad/tools/resp"
)

func Device(w http.ResponseWriter, r *http.Request) {
	var param dao.DeviceEntity

	err := DecryptToDao(r, &param) // 通用参数解析
	if err != nil {
		ErrorParam(w, err)
		return
	}

	// app_id, err := dao.GetAppByAppid(p["bundleid"])
	app_id, err := dao.GetAppByAppid(param.Bundleid)

	if err != nil || app_id == "" {
		ErrorParam(w, err)
		return
	}

	param.Appid = app_id
	param.Insert()

	Response(w, 200, "successful")
}
