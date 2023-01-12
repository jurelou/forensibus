package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	Splunk           SplunkConfiguration
	OutputFolder     string
	ArchivePasswords []string
}

type SplunkConfiguration struct {
	Address  string
	Port int
	Index string
}

func setDefaults() {
	viper.SetDefault("ArchivePasswords", []string{})
}

func Reload() error {
	// viper.SetDefault("ContentDir", "content")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Could not read config: %w", err)
	}
	err := viper.Unmarshal(&Config)
	if err != nil {
		return fmt.Errorf("Could not load config: %w", err)
	}
	return nil
}

func Configure() error {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./pipelines")
	viper.AddConfigPath("../pipelines")
	viper.SetConfigName("config")
	setDefaults()

	err := Reload()
	if err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return nil
}

func init() {
	if err := Configure(); err != nil {
		fmt.Println("Could not load config:", err)
	}
}
