package dao

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func Withdraw(app_id string, user_id string, money string, account string, realname string, task_type string, mobile string) (err error) {
	stmt, err := Db().Prepare("insert into `withdraw` (`app_id`,`user_id`,`money`,`account`,`realname`,`task_type`,`create_at`,`create_date`,`no`,`mobile`) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	timestame := time.Now().Unix()
	date := time.Now().Format("2006-01-02")
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	order := "x-" + time.Now().Format("20060102150304") + "-" + strconv.Itoa(rd.Intn(899)+100)
	_, err = stmt.Exec(app_id, user_id, money, account, realname, task_type, timestame, date, order, mobile)
	if err != nil {
		logrus.Error(err)
	}
	return
}

func WithdrawStatus(app_id string, user_id string) (id string, err error) {
	stmt, err := Db().Prepare("select id from withdraw where app_id=? and user_id=? and create_date=?")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()

	date := time.Now().Format("2006-01-02")

	row := stmt.QueryRow(app_id, user_id, date)

	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return id, nil
		}
		logrus.Error(err)
	}
	logrus.Info(id)
	return
}

func WithdrawCount(user_id string, task_type string) (count int, err error) {
	stmt, err := Db().Prepare("select count(id)as count_id from withdraw where user_id=? and task_type=? and create_date=?")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()

	date := time.Now().Format("2006-01-02")

	row := stmt.QueryRow(user_id, task_type, date)

	err = row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return count, nil
		}
		logrus.Error(err)
	}
	return
}
