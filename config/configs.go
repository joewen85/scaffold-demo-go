package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"scaffold-demo-go/utils/logs"
)

const (
	TimeFormat = "2006-01-02 15:04:05.000"
)

var Port string

func initLogSetting(logLevel string) {
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: TimeFormat})
}

func init() {
	defer logs.Info(nil, "初始化成功")
	// environment
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8080")
	viper.AutomaticEnv()

	// log setting
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	initLogSetting(logLevel)
}
