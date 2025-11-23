package config

import (
	"os"
	"user-service/common/util"
	"github.com/sirupsen/logrus"
)

var Config AppConfig

type AppConfig struct {
	Port                  int      `json:"port"`
	AppName               string   `json:"appName"`
	AppEnv                string   `json:"appEnv"`
	SignatureKey          string   `json:"SignatureKey"`
	Database              Database `json:"database"`
	RateLimiterMaxRequest float64  `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int      `json:"rateLimiterTimeSecond"`
	JwtSecretKey          string   `json:"jwtSecretKey"`
	JwtExpiredTime        int      `json:"jwtExpiredTime"`
}

type Database struct {
	Host               string `json:"host"`
	Port               int    `json:"port"`
	Name               string `json:"name"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	MaxOpenConnections int    `json:"maxOpenConnections"`
	MaxLifeConnections int    `json:"maxLifeConnections"`
	MaxIdleConnections int    `json:"maxIdleConnections"`
	MaxIdleTime        int    `json:"maxIdleTime"`
}

func Init() {
	err := util.BindFromJson(&Config, "config.json", ".")
	if err != nil {
		logrus.Infof("Failed to bind config: %v", err)
		err = util.BindFromConsul(&Config, os.Getenv("CONSUL_HTTP_URL"), os.Getenv("CONSUL_HTTP_KEY"))
		if err != nil {
			panic(err)
		}
	}
}