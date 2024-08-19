package router

import (
	"net/http"
)

func Hello(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func Hi(next func(http.ResponseWriter, map[string]interface{})) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := make(map[string]interface{})
		param["name"] = "lucifer"
		param["age"] = 29
		next(w, param)
		// fmt.Println(param) // 验证map是指针
	})
}
