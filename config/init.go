package config

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config AppConfig

func init() {
	initConfig()
	InitLog()
}

func initConfig() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("toml") // 按需设置
	_ = viper.ReadInConfig()

	err := viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}

	fmt.Println(Config)
}

// InitLog 初始化日志
func InitLog() {
	LogOutToFile := viper.GetBool("LogOutToFile")
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.SetReportCaller(true)
	if !LogOutToFile {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.TraceLevel)
	} else {
		filename := viper.GetString("LogDir") + "robot-" + time.Now().Format("2006-01-02") + ".log"
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		writers := []io.Writer{file, os.Stdout}
		fileAndConsole := io.MultiWriter(writers...)
		if err == nil {
			logrus.SetOutput(fileAndConsole)
		} else {
			logrus.Info("failed to logrus to file.")
		}
		logrus.SetLevel(logrus.InfoLevel)
	}
}
