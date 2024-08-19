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

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type LoginParam struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	// 参数解析
	var param LoginParam
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
	if param.Mobile == "" || param.Password == "" {
		ErrorParam(w, nil)
		return
	}
	logrus.Info("mobile:", param.Mobile)
	if len(param.Mobile) != 11 {
		ErrorParam(w, nil)
		return
	}

	// 查无此人
	user, err := dao.MobileUser(param.Mobile)
	if err != nil {
		logrus.Error(err)
		ErrorServer(w, err)
		return
	}
	logrus.Info(user)
	if user.Mobile == "" {
		Response(w, 400, "账号密码错误")
		return
	}

	// 密码错误
	pass := hash.Md5Check(param.Password, user.Password)
	if !pass {
		Response(w, 400, "账号密码错误")
		return
	}

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

func newToken(user dao.MobileUserEntity) (token string, err error) {

	// 登录成功
	uuid := uuid.NewV4()
	token = uuid.String()
	uid, err := rdbs.TokenGetUser(token)
	if err != nil {
		return "", err
	}
	if uid != "" {
		return newToken(user)
	}

	_, err = rdbs.TokenSet(token, user.Id)
	if err != nil {
		return "", err
	}
	return
}
