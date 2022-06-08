/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// routerCmd represents the router command
var routerCmd = &cobra.Command{
	Use:   "router",
	Short: "a cli tool to generate gin-framework router code.",
	Long:  "a cli tool to generate gin-framework router code.",
	Run: func(cmd *cobra.Command, args []string) {
		createRouter()
	},
}

func init() {
	rootCmd.AddCommand(routerCmd)

	// package name
	routerCmd.Flags().StringVarP(&data.PackageName, "package", "p", "", "gin-framework package name.")
	// name
	routerCmd.Flags().StringVarP(&data.Name, "name", "n", "", "gin-framework name.")
	// auth
	routerCmd.Flags().BoolVarP(&data.auth, "auth", "a", false, "gin-framework router is use auth middleware.")
}

func createRouter() {
	check()
	parseFile(createFile(ROUTER_FILE), "tpl/router.tpl")
}
