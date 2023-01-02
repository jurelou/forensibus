/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"errors"
	"github.com/spf13/cobra"
	"github.com/jurelou/forensibus/core"
	"github.com/jurelou/forensibus/utils"
)


var (
	pipelineconfig	string

	runCmd = &cobra.Command{
		Use:   "run [path]",
		Aliases: []string{"r"},
		Short: "A brief description of your command",
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
				if utils.FileExists(filepath) == false {
					return fmt.Errorf("%s is not a valid file or folder.\n", filepath)
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, filepaths []string) error {

			if utils.FileExists(pipelineconfig) == false {
				return fmt.Errorf("%s file does not exists.\n", pipelineconfig)
			}
			core.Yo(pipelineconfig, utils.UniqueListOfStrings(filepaths))
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&pipelineconfig, "pipeline", "p", "", "A pipeline configuration file.")
	runCmd.MarkPersistentFlagRequired("pipeline")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
