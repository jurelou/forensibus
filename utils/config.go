package utils

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"runtime"

// 	"github.com/fsnotify/fsnotify"
// 	"github.com/spf13/viper"
// )

// var Config Configuration

// type Configuration struct {
// 	TempFolder       string
// 	YaraRulesFolder  string
// 	SigmaRulesFolder string
// 	ProcessorTimeout int
// 	Splunk           SplunkConfiguration
// 	OutputFolder     string
// 	ArchivePasswords []string
// 	WorkersCount     uint32
// }

// type SplunkHECConfiguration struct {
// 	Address string // ok
// 	Token   string // ok
// }
// type SplunkConfiguration struct {
// 	// ManagementAddress  string
// 	Index string // ok
// 	Hec   SplunkHECConfiguration // ok
// }

// func setDefaults() {
// 	viper.SetDefault("ArchivePasswords", []string{})
// 	viper.SetDefault("WorkersCount", uint32(runtime.NumCPU()))
// }

// func bootstrap() error {
// 	return os.MkdirAll(Config.TempFolder, os.ModePerm)
// }

// func ReloadConfig() error {
// 	// viper.SetDefault("ContentDir", "content")
// 	if err := viper.ReadInConfig(); err != nil {
// 		return fmt.Errorf("could not read config: %w", err)
// 	}
// 	err := viper.Unmarshal(&Config)
// 	if err != nil {
// 		return fmt.Errorf("could not load config: %w", err)
// 	}

// 	return bootstrap()
// }

// func Configure() error {
// 	viper.AddConfigPath(".")
// 	viper.AddConfigPath("./config")
// 	viper.AddConfigPath("../config")
// 	viper.AddConfigPath("../../config")
// 	viper.AddConfigPath("../../../config")

// 	viper.SetConfigName("config")
// 	setDefaults()

// 	err := ReloadConfig()
// 	if err != nil {
// 		return err
// 	}

// 	viper.WatchConfig()
// 	viper.OnConfigChange(func(e fsnotify.Event) {
// 		err := ReloadConfig()
// 		if err != nil {
// 			fmt.Printf("Error while updating config: %s\n", err.Error())
// 			return
// 		}
// 		fmt.Println("Config file changed.")
// 		fmt.Printf("\t* OutputFolder: %s\n", Config.OutputFolder)
// 		fmt.Printf("\t* ArchivePasswords: %v\n", Config.ArchivePasswords)
// 		fmt.Printf("\t* ProcessorTimeout: %d\n", Config.ProcessorTimeout)
// 		fmt.Printf("\t* Splunk: %+v\n", Config.Splunk)
// 	})

// 	return nil
// }

// func init() {
// 	if err := Configure(); err != nil {
// 		log.Fatal(err.Error())
// 	}
// }
