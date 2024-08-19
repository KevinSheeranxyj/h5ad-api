package user

import (
	"net/http"

	"uy0/h5ad/server"
	. "uy0/h5ad/tools/resp"
)

type SignParam struct {
	Bundleid string `json:"bundleid"`
	UID      string `json:"uid"`
	UUID     string `json:"uuid"`
}

func SignIn(w http.ResponseWriter, param map[string]interface{}) {
	var err error

	appid := param["bundleid"].(string)
	uid := param["uid"].(string)
	uuid := param["uuid"].(string)

	if uuid == "" || uid == "" || appid == "" {
		ErrorParam(w, nil)
		return
	}

	err = server.SignIn(uuid, uid, appid)
	if err != nil {
		RespData(w, 400, "非法用户", nil)
		return
	}
	Response(w, 200, "sign successful")
	return
}

func SignStatus(w http.ResponseWriter, param map[string]interface{}) {
	uuid := param["uuid"].(string)

	if uuid != "" {
		status, err := server.SignStatus(uuid)
		if err != nil {
			ErrorParam(w, nil)
			return
		}
		data := map[string]string{"status": status}
		RespData(w, 200, "query successful", data)
		return
	}
	ErrorParam(w, nil)
}
