/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"errors"
	"github.com/spf13/cobra"
	afero "github.com/spf13/afero"
	"github.com/jurelou/forensibus/core"
)


var (
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
			appfs := afero.NewOsFs()
			for _, filepath := range args {
				_, err := appfs.Stat(filepath)
				if os.IsNotExist(err) {
					return fmt.Errorf("file %s does not exist.\n", filepath)
				}
			}
			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			core.Yo(args)
			
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}