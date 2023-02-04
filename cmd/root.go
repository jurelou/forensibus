package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "forensibus",
	Short: "...",
	Long: `The command-line interface allows to parse digital forensics artifacts.

There are two types of scanners:
- run: Processes each given file or folder, each artifact will be re-processed.
- watch: The application runs continuously, as soon as a new file is created it is processed.

To process artifacts, you need to have workers available, you can either:
1. Process artifacts locally (enabled by default when you execute "run" or "watch" command)
2. Start a new worker with the "worker" command.`,
}

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.forensibus.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
