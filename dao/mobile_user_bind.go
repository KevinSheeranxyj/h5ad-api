package dao

import (
	"github.com/sirupsen/logrus"
)

/**
 * 更新支付宝账号
 */
func Bind(user_id string, account string, realname string) (err error) {
	stmt, err := Db().Prepare("update `mobile_user` set `account`=?,`realname`=? where `id`=?")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(account, realname, user_id)
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}
