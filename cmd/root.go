/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	API_FILE = "api/[name]_api.go"
)

type Data struct {
	PackageName string
	Name        string
	ApiName     string
	Comment     string
	swagger     bool
}

func (d Data) IsUseSwagger() bool {
	return (d.Comment != "") && d.swagger
}

func (d Data) GetPackageName() string {
	if d.PackageName != "" {
		return d.PackageName
	}
	return "gin-framework"
}

var data Data

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gfctl",
	Short: "a cli tool to generate gin-framework api、service、model、router code.",
	Long:  `a cli tool to generate gin-framework api、service、model、router code.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
