package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	host         string
	port         int
	readTimeOut  int
	writeTimeOut int
}

func (sc ServerConfig) GetAddress() string {
	return fmt.Sprintf(":%d", sc.port)
}

func (sc ServerConfig) GetReadTimeOut() int {
	return sc.readTimeOut
}

func (sc ServerConfig) GetWriteTimeOut() int {
	return sc.writeTimeOut
}

func newServerConfig() ServerConfig {
	return ServerConfig{
		host:         viper.GetString("APP_HOST"),
		port:         viper.GetInt("APP_PORT"),
		readTimeOut:  viper.GetInt("APP_READ_TIMEOUT_IN_SEC"),
		writeTimeOut: viper.GetInt("APP_WRITE_TIMEOUT_IN_SEC"),
	}
}
