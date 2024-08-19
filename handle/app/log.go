package app

import (
	"net/http"

	"uy0/h5ad/dao"
	. "uy0/h5ad/tools/resp"
)

func Log(w http.ResponseWriter, r *http.Request) {
	var param dao.AdLogEntity

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	logrus.Info(err)
	// 	ErrorServer(w, err)
	// 	return
	// }
	// if err = json.Unmarshal(body, &param); err != nil {
	// 	ErrorParam(w, err)
	// 	return
	// }
	// fmt.Printf("%+v\n", param)
	// sign := hash.ByteMd5(body)
	// fmt.Println(sign)

	err := DecryptToDao(r, &param) // 通用参数解析
	if err != nil {
		ErrorParam(w, err)
		return
	}

	app_id, err := dao.GetAppByAppid(param.Bundleid)

	if err != nil || app_id == "" {
		ErrorParam(w, err)
		return
	}

	param.Appid = app_id
	param.Insert()

	Response(w, 200, "successful")
}
