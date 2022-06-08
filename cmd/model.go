/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "a cli tool to generate gin-framework model code.",
	Long:  "a cli tool to generate gin-framework model code.",
	Run: func(cmd *cobra.Command, args []string) {
		createModel()
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)

	// package name
	modelCmd.Flags().StringVarP(&data.PackageName, "package", "p", "", "gin-framework package name.")
	// name
	modelCmd.Flags().StringVarP(&data.Name, "name", "n", "", "gin-framework name.")
	// table name
	modelCmd.Flags().StringVarP(&data.TableName, "table", "t", "", "gin-framework table name.")
}

func createModel() {
	check()
	// 数据表名称
	if data.TableName == "" {
		fmt.Println("table name can't be empty.")
		os.Exit(1)
	}
	parseFile(createFile(MODEL_FILE), "tpl/model.tpl")
}
