package rdbs

import (
	"fmt"
	"testing"

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

func TestConfig(t *testing.T) {
	val, err := Config("name")
	fmt.Println(val, err)
}
