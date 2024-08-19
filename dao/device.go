package dao

import (
	"github.com/sirupsen/logrus"
)

type Dao interface {
	Insert() error
}

type DeviceEntity struct {
	Appid string `json:"app_id"`
	Uid   string `json:"uid"`
	Uuid  string `json:"uuid"`

	Bundleid       string `json:"bundleid"`
	PackageVersion string `json:"package_version"`

	Sys        string `json:"sys"`
	SysVersion string `json:"sys_version"`

	PhoneModel string `json:"phone_model"`
	MacAddr    string `json:"mac_addr"`

	SimStatus string `json:"sim_status"`
	// Md5 string `json:"md5"`
}

func (e *DeviceEntity) Insert() (err error) {
	stmt, err := Db().Prepare("insert into `device` (`app_id`,`uid`,`uuid`,`bundleid`,`package_version`,`sys`,`sys_version`,`phone_model`,`mac_addr`,`sim_status`) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logrus.Error(err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Appid, e.Uid, e.Uuid, e.Bundleid, e.PackageVersion, e.Sys, e.SysVersion, e.PhoneModel, e.MacAddr, e.SimStatus)
	if err != nil {
		logrus.Error(err)
	}
	return
}
