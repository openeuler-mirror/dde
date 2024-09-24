package config

import (
    "log"

    "github.com/spf13/viper"
)

type AlertConfig struct {
    Name   string                 `mapstructure:"name"`
    Type   string                 `mapstructure:"type"`
    Config map[string]interface{} `mapstructure:"configs"`
}

type Config struct {
    Paths        []string      `mapstructure:"path"`
    Alerts       []AlertConfig `mapstructure:"alert"`
    AIServiceURL string        `mapstructure:"ai_service_url"`
}

var GlobalConfig Config

func LoadConfig(configPath string) error {
    viper.SetConfigFile(configPath)
    if err := viper.ReadInConfig(); err != nil {
        return err
    }
    if err := viper.Unmarshal(&GlobalConfig); err != nil {
        return err
    }
    log.Printf("Configuration loaded: %+v\n", GlobalConfig)
    return nil
}
