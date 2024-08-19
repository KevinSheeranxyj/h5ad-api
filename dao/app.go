package dao

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

/**
 * 更新支付宝账号
 */
func App(name string) (app_id string, task_type string, amount string, err error) {
	stmt, err := Db().Prepare("select id,status,amount from app where name = ? and status !=0")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(name)

	err = row.Scan(&app_id, &task_type, &amount)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Error("empty app_id")
			return app_id, task_type, amount, nil
		}
		logrus.Error(err)
	}
	logrus.Info(app_id)
	return
}

/**
 * 更新支付宝账号
 */
func GetAppByAppid(appid string) (app_id string, err error) {
	stmt, err := Db().Prepare("select id from app where appid = ? and status !=0")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(appid)

	err = row.Scan(&app_id)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return app_id, nil
		// }
		logrus.Error(err)
		return
	}
	// logrus.Info(app_id)
	return
}
