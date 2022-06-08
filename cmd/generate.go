/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "a cli tool to generate gin-framework router|api|service|model code.",
	Long:  "a cli tool to generate gin-framework router|api|service|model code.",
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	// package name
	generateCmd.Flags().StringVarP(&data.PackageName, "package", "p", "", "gin-framework package name.")
	// api name
	generateCmd.Flags().StringVarP(&data.Name, "name", "n", "", "gin-framework name.")
	// api comment
	generateCmd.Flags().StringVarP(&data.Comment, "comment", "c", "", "gin-framework api comment.")
	// is use swagger
	generateCmd.Flags().BoolVarP(&data.swagger, "swagger", "s", false, "gin-framework api is use swagger.")
	// table name
	generateCmd.Flags().StringVarP(&data.TableName, "table", "t", "", "gin-framework table name.")
	// auth
	generateCmd.Flags().BoolVarP(&data.auth, "auth", "a", false, "gin-framework router is use auth middleware.")
}

func generate() {
	createRouter()
	createApi()
	createService()
	createModel()
}
