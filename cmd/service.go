/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "a cli tool to generate gin-framework service code.",
	Long:  "a cli tool to generate gin-framework service code.",
	Run: func(cmd *cobra.Command, args []string) {
		createService()
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	// package name
	serviceCmd.Flags().StringVarP(&data.PackageName, "package", "p", "", "gin-framework package name.")
	// name
	serviceCmd.Flags().StringVarP(&data.Name, "name", "n", "", "gin-framework name.")
}

func createService() {
	check()
	parseFile(createFile(SERVICE_FILE), "tpl/service.tpl")
}
