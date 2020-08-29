package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	host                  string
	port                  int
	name                  string
	user                  string
	password              string
	sslMode               string
	maxIdleConnections    int
	maxOpenConnections    int
	connectionMaxLifeTime int
}

func (dc DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dc.host, dc.port, dc.user, dc.password, dc.name, dc.sslMode)
}

func (dc DatabaseConfig) GetMaxIdleConnections() int {
	return dc.maxIdleConnections
}

func (dc DatabaseConfig) GetMaxOpenConnections() int {
	return dc.maxOpenConnections
}

func (dc DatabaseConfig) GetConnectionMaxLifeTime() int {
	return dc.connectionMaxLifeTime
}

func newDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		host:                  viper.GetString("DB_HOST"),
		port:                  viper.GetInt("DB_PORT"),
		name:                  viper.GetString("DB_NAME"),
		user:                  viper.GetString("DB_USER"),
		password:              viper.GetString("DB_PASSWORD"),
		sslMode:               viper.GetString("DB_SSL_MODE"),
		maxIdleConnections:    viper.GetInt("DB_MAX_IDLE_CONNECTIONS"),
		maxOpenConnections:    viper.GetInt("DB_MAX_OPEN_CONNECTIONS"),
		connectionMaxLifeTime: viper.GetInt("DB_CONNECTION_MAX_LIFETIME_IN_MIN"),
	}
}
