package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jurelou/forensibus/core/processors"
)

// processorCmd represents the processor command
var processorCmd = &cobra.Command{
	Use:     "processor [PROCESSOR_NAME]",
	Aliases: []string{"p", "proc"},
	Short:   "Get informations about available processors",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Args:  cobra.ExactArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 2 {
			return fmt.Errorf("requires at most ONE processor name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			processors.ListProcessors()
		} else {
			processors.RunSingleProcessor(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(processorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// processorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// processorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
