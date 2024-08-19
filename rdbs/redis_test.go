package rdbs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath("$HOME/go/src/robot/robot_wit/")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Info(err)
	}
}
