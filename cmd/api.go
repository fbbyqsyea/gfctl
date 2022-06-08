package cmd

import (
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
	apiCmd.Flags().StringVarP(&data.Name, "name", "n", "", "gin-framework name.")
	// api comment
	apiCmd.Flags().StringVarP(&data.Comment, "comment", "c", "", "gin-framework api comment.")
	// is use swagger
	apiCmd.Flags().BoolVarP(&data.swagger, "swagger", "s", false, "gin-framework api is use swagger.")

}

func createApi() {
	check()
	f := createFile(API_FILE)
	parseFile(f, "tpl/api.tpl")
}
