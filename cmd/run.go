/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	run "github.com/jurelou/forensibus/core/run"
	"github.com/jurelou/forensibus/utils"
	"github.com/spf13/cobra"
)

var (
	pipelineconfig string

	runCmd = &cobra.Command{
		Use:     "run [path]",
		Aliases: []string{"r"},
		Short:   "A brief description of your command",
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
					return fmt.Errorf("%s is not a valid file or folder.\n", filepath)
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, filepaths []string) error {
			if !utils.FileExists(pipelineconfig) {
				return fmt.Errorf("%s file does not exists.\n", pipelineconfig)
			}
			run.Run(pipelineconfig, utils.UniqueListOfStrings(filepaths))
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&pipelineconfig, "pipeline", "p", "", "A pipeline configuration file.")
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
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
