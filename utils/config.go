package utils

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	Splunk SplunkConfiguration
}

type SplunkConfiguration struct {
	Address  string
	Username string
	Password string
}

func Reload() error {
	//viper.SetDefault("ContentDir", "content")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		return err //log.Fatalf("unable to decode into struct, %v", err)
	}
	return nil
}
func Configure() error {
	// viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./pipelines")
	viper.AddConfigPath("../pipelines")
	viper.SetConfigName("config")

	err := Reload(); if err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return nil
}
