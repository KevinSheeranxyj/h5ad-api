package router

import (
	"net/http"
	"uy0/h5ad/handle/alipay"
	"uy0/h5ad/handle/app"
	"uy0/h5ad/handle/app/common"
	"uy0/h5ad/handle/common/sms"
	"uy0/h5ad/handle/common/welcome"
	"uy0/h5ad/handle/h5/user"
	"uy0/h5ad/handle/h5/withdraw"
)

func Http() {
	r := NewRouter()
	r.Use(cors)

	r.Add("/reg", user.Jump)

	r.Add("/api/common/config", common.Config)

	r.Add("/api/sms/reg", sms.Reg)
	r.Add("/api/user/reg", user.Reg)
	r.Add("/api/user/login", user.Login)
	r.Add("/api/alipay/bind_status", alipay.BindStatus)
	r.Add("/api/withdraw/bonus", withdraw.Bonus)

	r.Add("/api/user/signin", Final(user.SignIn))
	r.Add("/api/user/signstatus", Final(user.SignStatus))

	r.Add("/api/device/device", app.Device)
	r.Add("/api/ad/log", app.Log)

	r.Add("/hi", Final(welcome.Hi))

	r.Load()

	http.HandleFunc("/", Hello(welcome.Welcome))
	// http.HandleFunc("/hi", cors(Final(welcome.Hi)))
}
