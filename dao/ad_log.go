package dao

import (
	"time"

	"github.com/sirupsen/logrus"
)

type AdLogEntity struct {
	Appid string `json:"app_id"`
	Uid   string `json:"uid"`
	Uuid  string `json:"uuid"`

	Bundleid string `json:"bundleid"`

	PlatformName string `json:"platform_name"`
	AdLocal      string `json:"ad_local"`
}

func (e *AdLogEntity) Insert() (err error) {
	stmt, err := Db().Prepare("insert into `ad_log` (`app_id`,`uid`,`uuid`,`bundleid`,`platform_name`,`ad_local`,`date`) values(?,?,?,?,?,?,?)")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	date := time.Now().Format("2006-01-02")
	_, err = stmt.Exec(e.Appid, e.Uid, e.Uuid, e.Bundleid, e.PlatformName, e.AdLocal, date)
	if err != nil {
		logrus.Error(err)
	}
	return
}
