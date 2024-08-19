package main

import (
	"fmt"
	"net/http"

	"uy0/h5ad/config"
	"uy0/h5ad/router"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("listen port", config.Config.App.Port)
	fmt.Println("-----------------------------------------------")

	router.Http()
	err := http.ListenAndServe(config.Config.App.Port, nil)

	if err != nil {
		fmt.Println("Server error")
	}
}
