package cmd

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	run "github.com/jurelou/forensibus/core/run"
	"github.com/jurelou/forensibus/utils"
)

var (
	// pipelineconfig string
	// tag            string
	// disableWorker  bool
	// verbose        bool
	runArgs = run.NewRunargs()

	runCmd = &cobra.Command{
		Use:     "run [path]",
		Aliases: []string{"r"},
		Short:   "Run a pipeline against a given file or folder",
		Long: `
Examples:
- Parse a whole folder containing DFIR-ORC archives:
$ forensibus run -p pipelines/dfir-orc.hcl /data/MY_ORCS

- Upload results to a custom splunk index (index should already exists)
$ forensibus run -p pipelines/dfir-orc.hcl /data/MY_ORCS --splunk_index myCustomIndex

- Upload results to a custom splunk instance
$ forensibus run -p pipelines/dfir-orc.hcl /data/MY_ORCS --splunk_address http://splunk:8088 --splunk_token <HEC_TOKEN>
		`,
		// Args: func(cmd *cobra.Command, args []string) error {
		// 	if len(args) < 1 {
		// 		return errors.New("requires a filepath  argument")
		// 	}
		// 	for _, filepath := range args {
		// 		if !utils.FileExists(filepath) {
		// 			return fmt.Errorf("%s is not a valid file or folder", filepath)
		// 		}
		// 	}
		// 	return nil
		// },

		RunE: func(cmd *cobra.Command, filepaths []string) error {
			runArgs.Targets = utils.UniqueListOfStrings(filepaths)
			if err := runArgs.Validate(); err != nil {
				pterm.Error.Println(err.Error())
				return nil
			}
			run.Run(runArgs)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&runArgs.PipelineFile, "pipeline", "p", "", "A pipeline configuration file.")
	runCmd.PersistentFlags().StringVarP(&runArgs.Tag, "tag", "t", "", "Tag the process to identify it more easily (defaults to a randomly generated string)")
	runCmd.PersistentFlags().BoolVarP(&runArgs.DisableLocalWorker, "disable_worker", "d", false, "Disable local worker (defaults to: false)")
	runCmd.PersistentFlags().BoolVarP(&runArgs.Verbose, "verbose", "v", false, "Increase logs verbosity (defaults to: false)")

	runCmd.PersistentFlags().StringVarP(&runArgs.Splunk.Index, "splunk_index", "i", "main", "Splunk index to use (defaults to: main)")
	runCmd.PersistentFlags().StringVarP(&runArgs.Splunk.Token, "splunk_token", "s", "42424242-4242-4242-4242-424242424242", "Splunk HEC token to use (defaults to: 42424242-4242-4242-4242-424242424242)")
	runCmd.PersistentFlags().StringVarP(&runArgs.Splunk.Address, "splunk_address", "a", "http://localhost:8088", "Splunk index to use (defaults to: http://localhost:8088)")

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
