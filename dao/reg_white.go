package dao

import (
	"github.com/sirupsen/logrus"
)

func RegWhite(mobile string) (id int, err error) {
	stmt, err := Db().Prepare("select id from reg_white_list where mobile=? and status=1")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	res := stmt.QueryRow(mobile)

	err = res.Scan(&id)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(id)
	return
}
