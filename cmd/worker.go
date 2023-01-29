/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/jurelou/forensibus/worker"
)

// workerCmd represents the worker command
var (
	workersCount uint32

	workerCmd = &cobra.Command{
		Use:     "worker",
		Aliases: []string{"w"},
		Short:   "Run a processing worker",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			worker.Run(workersCount)
		},
	}
)

func init() {
	workerCmd.PersistentFlags().Uint32VarP(&workersCount, "workers", "w", uint32(runtime.NumCPU()), "Number of workers to launch (defaults to the number of CPUs).")

	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
