package cmd

import (
	"errors"
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"github.com/jurelou/forensibus/core/run"
	"github.com/jurelou/forensibus/core/watch"
	"github.com/jurelou/forensibus/utils"
)

var (
	watchArgs = run.NewRunargs()

	watchCmd = &cobra.Command{
		Use:     "watch [path]",
		Aliases: []string{"r"},
		Short:   "Watch a given folder, fun a pipeline against each new file",
		Long: `Each newly created file will be scanned
	Existing files will not be scanned.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a filepath  argument")
			}
			for _, filepath := range args {
				if !utils.FileExists(filepath) {
					return fmt.Errorf("%s is not a valid file or folder", filepath)
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, filepaths []string) error {
			watchArgs.Targets = utils.UniqueListOfStrings(filepaths)
			if err := watchArgs.Validate(); err != nil {
				pterm.Error.Println(err.Error())
				return nil
			}
			watch.Watch(watchArgs)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.PersistentFlags().StringVarP(&watchArgs.PipelineFile, "pipeline", "p", "", "A pipeline configuration file.")
	watchCmd.PersistentFlags().StringVarP(&watchArgs.Tag, "tag", "t", "", "Tag the process to identify it more easily (defaults to a randomly generated string)")
	watchCmd.PersistentFlags().BoolVarP(&watchArgs.DisableLocalWorker, "disable_worker", "d", false, "Disable local worker (defaults to: false)")
	watchCmd.PersistentFlags().BoolVarP(&watchArgs.Verbose, "verbose", "v", false, "Increase logs verbosity (defaults to: false)")

	watchCmd.PersistentFlags().StringVarP(&watchArgs.Splunk.Index, "splunk_index", "i", "main", "Splunk index to use (defaults to: main)")
	watchCmd.PersistentFlags().StringVarP(&watchArgs.Splunk.Token, "splunk_token", "s", "42424242-4242-4242-4242-424242424242", "Splunk HEC token to use (defaults to: 42424242-4242-4242-4242-424242424242)")
	watchCmd.PersistentFlags().StringVarP(&watchArgs.Splunk.Address, "splunk_address", "a", "http://localhost:8088", "Splunk index to use (defaults to: http://localhost:8088)")

	err := watchCmd.MarkPersistentFlagRequired("pipeline")
	if err != nil {
		fmt.Printf("Could nork mark pipeline as a persistent flag")
	}
}
