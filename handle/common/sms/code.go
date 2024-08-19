package sms

import (
	"net/http"
	"strings"
	"uy0/h5ad/dao"
	"uy0/h5ad/rdbs"
	. "uy0/h5ad/tools/resp"

	"github.com/sirupsen/logrus"
)

func Reg(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	mobile := query.Get("mobile")

	if mobile == "" {
		logrus.Error("empty mobile")
		ErrorParam(w, nil)
		return
	}
	mobile = strings.Replace(mobile, " ", "", -1)
	if mobile == "" {
		logrus.Error("empty mobile")
		ErrorParam(w, nil)
		return
	}
	logrus.Info("mobile:", mobile)

	if len(mobile) != 11 {
		logrus.Error("mobile lenth err")
		ErrorParam(w, nil)
		return
	}

	res, err := rdbs.RedisLock(mobile, "1", 60)
	if err != nil {
		logrus.Error(err)
		return
	}
	if !res {
		Response(w, 200, "请60秒后再试")
		return
	}

	reg_white, err := rdbs.AppConfig("reg_white_list")
	if err != nil {
		ErrorServer(w, err)
		return
	}
	logrus.Info("reg_white:", reg_white)

	if reg_white == "1" {
		id, err := dao.RegWhite(mobile)
		if err != nil || id == 0 {
			Response(w, 400, "当前手机号不在白名单内")
			return
		}
	}
	err = rdbs.SmsSend(mobile)
	if err != nil {
		ErrorServer(w, err)
		return
	}
	Response(w, 200, "发送成功")
	return
}
