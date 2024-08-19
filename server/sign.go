package server

import (
	"uy0/h5ad/dao"
	"uy0/h5ad/rdbs"
)

func SignIn(uuid string, uid string, appid string) (err error) {
	status, err := SignStatus(uuid)
	if status == "1" {
		return
	}
	err = dao.SignIn(uuid, uid, appid)
	if err != nil {
		return
	}
	err = rdbs.SignIn(uuid)
	return
}

func SignStatus(uuid string) (status string, err error) {
	status, err = rdbs.SignStatus(uuid)
	return
}
