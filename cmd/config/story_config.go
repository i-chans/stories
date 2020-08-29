package config

import "github.com/spf13/viper"

type StoryConfig struct {
	titleMaxLength int
	bodyMaxLength  int
}

func (sc StoryConfig) GetTitleMaxLength() int {
	return sc.titleMaxLength
}

func (sc StoryConfig) GetBodyMaxLength() int {
	return sc.bodyMaxLength
}

func newStoryConfig() StoryConfig {
	return StoryConfig{
		titleMaxLength: viper.GetInt("TITLE_MAX_LENGTH"),
		bodyMaxLength:  viper.GetInt("BODY_MAX_LENGTH"),
	}
}
