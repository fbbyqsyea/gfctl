package cmd

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/fbbyqsyea/gfctl/util"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "a cli tool to generate gin-framework api code.",
	Long:  "a cli tool to generate gin-framework api code.",
	Run: func(cmd *cobra.Command, args []string) {
		createApi()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	// package name
	apiCmd.Flags().StringVarP(&data.PackageName, "package", "p", "", "gin-framework package name.")
	// api name
	apiCmd.Flags().StringVarP(&data.Name, "name", "n", "", "gin-framework api name.")
	// api comment
	apiCmd.Flags().StringVarP(&data.Comment, "comment", "c", "", "gin-framework api comment.")
	// is use swagger
	apiCmd.Flags().BoolVarP(&data.swagger, "swagger", "s", false, "gin-framework api is use swagger.")

}

func createApi() {
	// check api name
	if data.Name == "" {
		fmt.Println("api name can't be empty.")
		os.Exit(1)
	}
	data.ApiName = util.FirstUpper(data.Name)
	apiFile := strings.ReplaceAll(API_FILE, "[name]", data.Name)
	os.MkdirAll(filepath.Dir(apiFile), fs.FileMode(os.O_CREATE))
	os.Chmod(filepath.Dir(apiFile), 0777)
	f, err := os.OpenFile(apiFile, os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Printf("create api file failed, %v\n", err)
		os.Exit(1)
	}

	tpl, err := template.ParseFiles("tpl/api.tpl")
	if err != nil {
		fmt.Printf("parse api file failed, %v\n", err)
		os.Exit(1)
	}
	//渲染模板
	err = tpl.Execute(f, data)
	if err != nil {
		fmt.Printf("render api file failed, %v\n", err)
		os.Exit(1)
	}
}
