package dao

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

func SignIn(uuid string, uid string, appid string) (err error) {
	stmt, err := Db().Prepare("insert into `sign_record` (`uuid`,`uid`,`app_id`,`date`) values(?,?,?,?)")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	date := time.Now().Format("2006-01-02")

	app_id, err := GetAppByAppid(appid)
	if err != nil || app_id == "" {
		return
	}

	_, err = stmt.Exec(uuid, uid, app_id, date)
	if err != nil {
		logrus.Error(err)
	}
	return
}

func SignStatus(app_id string, uid string) (id string, err error) {
	stmt, err := Db().Prepare("select id from sign_record where app_id=? and uid=? and date=?")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()

	date := time.Now().Format("2006-01-02")

	row := stmt.QueryRow(app_id, uid, date)

	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			logrus.Error("empty")
			return id, nil
		}
		logrus.Error(err)
	}
	logrus.Info(id)
	return
}
