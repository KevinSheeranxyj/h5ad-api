package user

import (
	"net/http"
	"strings"
	"uy0/h5ad/config"
)

// Jump 跳转
func Jump(w http.ResponseWriter, r *http.Request) {
	url := config.Config.App.Url
	if url == "" {
		url = "https://" + r.Host
	}
	if strings.Contains(r.UserAgent(), "MicroMessenger") {
		http.Redirect(w, r, url+"/h5/user/reg.html", http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, url+"/h5/tips/wechat.html", http.StatusTemporaryRedirect)
	}
}
