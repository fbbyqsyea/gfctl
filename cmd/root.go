/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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

const (
	API_FILE     = "api/[name]_api.go"
	SERVICE_FILE = "service/[name]_service.go"
	MODEL_FILE   = "model/[name]_model.go"
	ROUTER_FILE  = "router/[name]_api_router.go"
)

type Data struct {
	PackageName string
	Name        string
	VarName     string
	Comment     string
	swagger     bool
	Host        string
	UserName    string
	Password    string
	DbName      string
	TableName   string
	auth        bool
}

func (d Data) IsUseSwagger() bool {
	return (d.Comment != "") && d.swagger
}

func (d Data) IsUseAuth() bool {
	return d.auth
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
	Short: "a cli tool to generate gin-framework and auto create router|api|contexts|service|model code.",
	Long:  `a cli tool to generate gin-framework and auto create router|api|contexts|service|model code.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// 必须参数校验
func check() {
	// check api name
	if data.Name == "" {
		fmt.Println("name can't be empty.")
		os.Exit(1)
	}
	data.VarName = util.FirstUpper(data.Name)
}

func createFile(tpl string) *os.File {
	file := strings.ReplaceAll(tpl, "[name]", data.Name)
	os.MkdirAll(filepath.Dir(file), fs.FileMode(os.O_CREATE))
	os.Chmod(filepath.Dir(file), 0777)
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		fmt.Printf("create file failed, %v\n", err)
		os.Exit(1)
	}
	return f
}

func parseFile(f *os.File, tplPath string) {
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		fmt.Printf("parse file failed, %v\n", err)
		os.Exit(1)
	}
	err = tpl.Execute(f, data)
	if err != nil {
		fmt.Printf("render file failed, %v\n", err)
		os.Exit(1)
	}
}
