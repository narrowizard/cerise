package config

import (
	"github.com/narrowizard/cerise/models"
	"github.com/spf13/viper"
)

// LoadConfig read config.yaml and resolve to models.Task array
func LoadConfig() (map[string]models.Task, error) {
	var v = viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	var err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var m = make(map[string]models.Task)
	err = v.UnmarshalKey("tasks", &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
