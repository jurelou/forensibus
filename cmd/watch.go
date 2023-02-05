package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jurelou/forensibus/core/watch"
	"github.com/jurelou/forensibus/utils"
)

var (
	watchPipelineConfig string
	watchTag            string
	watchDisableWorker  bool
	watchVerbose        bool

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
			if !utils.FileExists(watchPipelineConfig) {
				return fmt.Errorf("%s file does not exists", watchPipelineConfig)
			}
			watch.Watch(watchPipelineConfig, utils.UniqueListOfStrings(filepaths), watchTag, watchDisableWorker, verbose)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(watchCmd)

	watchCmd.PersistentFlags().StringVarP(&watchPipelineConfig, "pipeline", "p", "", "A pipeline configuration file.")
	watchCmd.PersistentFlags().StringVarP(&watchTag, "tag", "t", "", "Tag the process to identify it more easily")
	watchCmd.PersistentFlags().BoolVarP(&watchDisableWorker, "disable_worker", "d", false, "Disable local worker")
	watchCmd.PersistentFlags().BoolVarP(&watchVerbose, "verbose", "v", false, "Increase logs verbosity")

	err := watchCmd.MarkPersistentFlagRequired("pipeline")
	if err != nil {
		fmt.Printf("Could nork mark pipeline as a persistent flag")
	}
}
