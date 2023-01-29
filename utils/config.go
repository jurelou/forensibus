package utils

import (
	"fmt"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	YaraRulesFolder	string
	SigmaRulesFolder string
	ProcessorTimeout int
	Splunk           SplunkConfiguration
	OutputFolder     string
	ArchivePasswords []string
	WorkersCount     uint32
}

type SplunkHECConfiguration struct {
	Address string
	Token   string
}
type SplunkConfiguration struct {
	// ManagementAddress  string
	Index string
	Hec   SplunkHECConfiguration
}

func setDefaults() {
	viper.SetDefault("ArchivePasswords", []string{})
	viper.SetDefault("WorkersCount", uint32(runtime.NumCPU()))
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
	viper.AddConfigPath("../../pipelines")
	viper.SetConfigName("config")
	setDefaults()

	err := Reload()
	if err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := Reload()
		if err != nil {
			fmt.Printf("Error while updating config: %s\n", err.Error())
			return
		}
		fmt.Println("Config file changed.")
		fmt.Printf("\t* OutputFolder: %s\n", Config.OutputFolder)
		fmt.Printf("\t* ArchivePasswords: %v\n", Config.ArchivePasswords)
		fmt.Printf("\t* ProcessorTimeout: %d\n", Config.ProcessorTimeout)
		fmt.Printf("\t* Splunk: %+v\n", Config.Splunk)
	})

	return nil
}

func init() {
	if err := Configure(); err != nil {
		fmt.Println("Could not load config:", err)
	}
}
