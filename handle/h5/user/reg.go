package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"uy0/h5ad/dao"
	"uy0/h5ad/rdbs"
	"uy0/h5ad/tools/hash"
	. "uy0/h5ad/tools/resp"

	"github.com/sirupsen/logrus"
)

type RegParam struct {
	Mobile         string `json:"mobile"`
	Password       string `json:"password"`
	Code           string `json:"code"`
	VerifyPassword string `json:"verify_passowrd"`
}

func Reg(w http.ResponseWriter, r *http.Request) {

	// 参数解析
	var param RegParam
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		ErrorServer(w, err)
		return
	}
	logrus.Info(body)
	if err = json.Unmarshal(body, &param); err != nil {
		logrus.Error(err)
		ErrorParam(w, err)
		return
	}
	fmt.Printf("%+v\n", param)

	// 手机号空
	param.Mobile = strings.Replace(param.Mobile, " ", "", -1)
	if param.Mobile == "" || param.Password == "" || param.Password != param.VerifyPassword {
		ErrorParam(w, nil)
		return
	}
	logrus.Info("mobile:", param.Mobile)
	if len(param.Mobile) != 11 {
		ErrorParam(w, nil)
		return
	}

	code, err := rdbs.SmsGetCode(param.Mobile)
	if err != nil {
		ErrorServer(w, err)
		return
	}

	if code == "" {
		Response(w, 400, "验证码错误")
		return
	}
	// 验证码
	if param.Code != code {
		Response(w, 400, "验证码错误")
		return
	}

	// 查有此人
	user, err := dao.MobileUser(param.Mobile)
	if err != nil {
		logrus.Error(err)
		ErrorServer(w, err)
		return
	} else {
		logrus.Info(user)
		if user.Mobile != "" {
			Response(w, 400, "已经注册，请直接登录")
			return
		}
	}

	// 注册
	user, err = dao.Reg(param.Mobile, hash.Md5(param.Password))
	if err != nil {
		ErrorServer(w, err)
		return
	}
	logrus.Info(user)

	token, err := newToken(user)
	if err != nil {
		ErrorServer(w, err)
		return
	}
	day := time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	data := map[string]string{
		"token":       token,
		"period":      "86400",
		"expire_time": day,
	}

	RespData(w, 200, "successful", data)
}
