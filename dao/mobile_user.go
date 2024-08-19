package dao

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type MobileUserEntity struct {
	Id       int
	Mobile   string
	Password string
	Account  string
	Realname string
}

/**
 * 查询用户
 * 登录 查支付宝绑定状态
 */
func MobileUserById(id string) (user MobileUserEntity, err error) {
	stmt, err := Db().Prepare("select id,account,realname,mobile from mobile_user where id=? and status=1")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	err = row.Scan(&user.Id, &user.Account, &user.Realname, &user.Mobile)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		logrus.Error(err)
	}
	logrus.Info(user)
	return
}

/**
 * 查询用户
 * 登录 查支付宝绑定状态
 */
func MobileUser(mobile string) (user MobileUserEntity, err error) {
	stmt, err := Db().Prepare("select id,mobile,password from mobile_user where mobile=? and status=1")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(mobile)

	err = row.Scan(&user.Id, &user.Mobile, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		logrus.Error(err)
	}
	logrus.Info(user)
	return
}

/**
 * 注册
 * 先插入 后查询
 */
func Reg(mobile string, password string) (user MobileUserEntity, err error) {
	stmt, err := Db().Prepare("insert into mobile_user (mobile,password) values (?,?)")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(mobile, password)
	if err != nil {
		logrus.Error(err)
		return
	}

	stmtQuery, err := Db().Prepare("select id,mobile,password from mobile_user where mobile=? and status=1")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	row := stmtQuery.QueryRow(mobile)

	err = row.Scan(&user.Id, &user.Mobile, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		logrus.Error(err)
	}
	logrus.Info(user)
	return
}
