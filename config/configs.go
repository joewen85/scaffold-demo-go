package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"scaffold-demo-go/utils/logs"
)

const (
	TimeFormat = "2006-01-02 15:04:05.000"
)

var (
	Port       string
	SigningKey string
	ExpireTime uint64
	Username   string
	Password   string
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func initLogSetting(logLevel string) {
	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
		_ = os.Setenv("SECRET_KEY", "joe123456")
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

	//	jwt
	SigningKey = viper.GetString("SECRET_KEY")
	if SigningKey == "" {
		logs.Error(nil, "environment not set SECRET_KEY")
		os.Exit(1)
	}
	viper.SetDefault("JWT_EXPIRE_TIME", 86400)
	ExpireTime = viper.GetUint64("JWT_EXPIRE_TIME")
}
