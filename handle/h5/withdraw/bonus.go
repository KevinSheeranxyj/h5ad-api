package withdraw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"uy0/h5ad/dao"
	"uy0/h5ad/rdbs"
	. "uy0/h5ad/tools/resp"

	"github.com/sirupsen/logrus"
)

type BonusParam struct {
	Account  string
	Realname string

	Appname string
	Uid     string
}

func Bonus(w http.ResponseWriter, r *http.Request) {

	// token参数

	query := r.URL.Query()
	token := query.Get("token")
	if token == "" {
		ErrorAuth(w)
		return
	}

	// post参数解析
	var param BonusParam
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err)
		ErrorServer(w, err)
		return
	}
	if err = json.Unmarshal(body, &param); err != nil {
		logrus.Error(err)
		ErrorParam(w, err)
		return
	}
	fmt.Printf("%+v\n", param)

	// 用户信息
	user_id, err := rdbs.TokenGetUser(token)
	if err != nil {
		ErrorServer(w, nil)
		return
	}
	if user_id == "" {
		ErrorAuth(w)
		return
	}
	user, err := dao.MobileUserById(user_id)
	if err != nil {
		ErrorServer(w, nil)
		return
	}

	// 绑定支付宝
	if user.Account == "" {
		if param.Account == "" || param.Realname == "" {
			ErrorParam(w, nil)
			return
		}
		user.Account = param.Account
		user.Realname = param.Realname

		dao.Bind(user_id, param.Account, param.Realname)
	}
	// 判断签到
	// 软件名查id
	app_id, task_type, amount, err := dao.App(param.Appname)
	if err != nil {
		ErrorServer(w, nil)
		return
	}
	if app_id == "" {
		ErrorParam(w, err)
		return
	}

	sign_id, err := dao.SignStatus(app_id, param.Uid)
	if err != nil {
		ErrorServer(w, nil)
		return
	}
	if sign_id == "" {
		Response(w, 400, "未进行签到")
		return
	}

	// 判断今日是否已提现
	withdraw_id, err := dao.WithdrawStatus(app_id, user_id)
	if err != nil {
		ErrorServer(w, nil)
		return
	}
	if withdraw_id != "" {
		Response(w, 400, "今日已领取")
		return
	}

	// 判断每日任务提现次数
	count, err := dao.WithdrawCount(user_id, task_type)
	if err != nil {
		ErrorServer(w, err)
		return
	}

	if task_type == "1" {
		if count >= 2 {
			Response(w, 400, "今日已达领红包上限")
			return
		}
	} else {
		if count >= 5 {
			Response(w, 400, "今日已达领红包上限")
			return
		}
	}
	dao.Withdraw(app_id, user_id, amount, user.Account, user.Realname, task_type, user.Mobile)

	Response(w, 200, "申请成功，请等待到账")
}
