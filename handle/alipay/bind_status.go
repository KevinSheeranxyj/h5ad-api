package alipay

import (
	"net/http"
	"uy0/h5ad/dao"
	"uy0/h5ad/rdbs"
	. "uy0/h5ad/tools/resp"
)

func BindStatus(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	token := query.Get("token")
	if token == "" {
		ErrorAuth(w)
		return
	}

	uid, err := rdbs.TokenGetUser(token)
	if err != nil {
		ErrorServer(w, nil)
		return
	}
	if uid == "" {
		ErrorAuth(w)
		return
	}

	user, err := dao.MobileUserById(uid)
	if err != nil {
		ErrorServer(w, nil)
		return
	}

	var status string

	if user.Account == "" {
		status = "0"
	} else {
		status = "1"
	}

	data := map[string]string{
		"status": status,
	}
	RespData(w, 200, "请求成功", data)
}
