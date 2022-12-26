/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "scli",
	Short: "scli is a application to interact with SberCloud via CLI",
	Long:  `scli is a application to interact with SberCloud via CLI.`,
}

var jmesPathQuery string

func Execute() {
	ProjectID = os.Getenv("PROJECT_ID")

	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
