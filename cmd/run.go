package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	run "github.com/jurelou/forensibus/core/run"
	"github.com/jurelou/forensibus/utils"
)

var (
	pipelineconfig string
	tag            string
	disableWorker  bool
	verbose bool

	runCmd = &cobra.Command{
		Use:     "run [path]",
		Aliases: []string{"r"},
		Short:   "Run a pipeline against a given file or folder",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
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
			if !utils.FileExists(pipelineconfig) {
				return fmt.Errorf("%s file does not exists", pipelineconfig)
			}
			run.Run(pipelineconfig, utils.UniqueListOfStrings(filepaths), tag, disableWorker, verbose)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&pipelineconfig, "pipeline", "p", "", "A pipeline configuration file.")
	runCmd.PersistentFlags().StringVarP(&tag, "tag", "t", "", "Tag the process to identify it more easily")
	runCmd.PersistentFlags().BoolVarP(&disableWorker, "disable_worker", "d", false, "Disable local worker")
	runCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Increase logs verbosity")

	// runCmd.PersistentFlags().StringVarP(&splunkIndex, "splunk-index", "s", "", "Change default splunk index (WON'T BE CREATED).")

	err := runCmd.MarkPersistentFlagRequired("pipeline")
	if err != nil {
		fmt.Printf("Could nork mark pipeline as a persistent flag")
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
