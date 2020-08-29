package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	env            string
	migrationPath  string
	databaseConfig DatabaseConfig
	serverConfig   ServerConfig
	storyConfig    StoryConfig
}

func (c Config) GetEnv() string {
	return c.env
}

func (c Config) GetDatabaseConfig() DatabaseConfig {
	return c.databaseConfig
}

func (c Config) GetServerConfig() ServerConfig {
	return c.serverConfig
}

func (c Config) GetStoryConfig() StoryConfig {
	return c.storyConfig
}

func (c Config) GetMigrationPath() string {
	return c.migrationPath
}

func NewConfig() Config {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	return Config{
		env:            viper.GetString("ENV"),
		migrationPath:  viper.GetString("MIGRATION_PATH"),
		databaseConfig: newDatabaseConfig(),
		serverConfig:   newServerConfig(),
		storyConfig:    newStoryConfig(),
	}
}
