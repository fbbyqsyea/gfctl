/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// contextsCmd represents the contexts command
var contextsCmd = &cobra.Command{
	Use:   "contexts",
	Short: "a cli tool to generate gin-framework contexts code.",
	Long:  "a cli tool to generate gin-framework contexts code.",
	Run: func(cmd *cobra.Command, args []string) {
		createContexts()
	},
}

// // rows, err := MtsqlDb.Query(`select COLUMN_NAME ,DATA_TYPE,IFNULL(CHARACTER_MAXIMUM_LENGTH,0),COLUMN_TYPE,IFNULL(NUMERIC_PRECISION,0),IFNULL(NUMERIC_SCALE,0),COLUMN_COMMENT,column_key,extra,ORDINAL_POSITION  from information_schema.COLUMNS where  table_schema = ? and  table_name = ?`, dbname, tablename)

func init() {
	rootCmd.AddCommand(contextsCmd)

	// package name
	contextsCmd.Flags().StringVarP(&data.PackageName, "package", "p", "", "gin-framework package name.")
	// name
	contextsCmd.Flags().StringVarP(&data.Name, "name", "n", "", "gin-framework name.")
	// host
	contextsCmd.Flags().StringVarP(&data.Host, "host", "h", "", "gin-framework db host.")
	// user name
	contextsCmd.Flags().StringVarP(&data.UserName, "user", "u", "", "gin-framework db user.")
	// password
	contextsCmd.Flags().StringVarP(&data.Password, "password", "p", "", "gin-framework db password.")
	// db name
	contextsCmd.Flags().StringVarP(&data.DbName, "database", "d", "", "gin-framework db name.")
	// table name
	contextsCmd.Flags().StringVarP(&data.TableName, "table", "t", "", "gin-framework table name.")
}

func createContexts() {
	check()
	if data.Host == "" {
		fmt.Println("db host can't by empty.")
		os.Exit(1)
	}
	if data.UserName == "" {
		fmt.Println("db user can't by empty.")
		os.Exit(1)
	}
	if data.Password == "" {
		fmt.Println("db password can't by empty.")
		os.Exit(1)
	}
	if data.DbName == "" {
		fmt.Println("db name can't by empty.")
		os.Exit(1)
	}
	if data.TableName == "" {
		fmt.Println("db table name can't by empty.")
		os.Exit(1)
	}
}

// type Column struct {
// 	ColumnName      string        `json:"colname"`
// 	ColumnJsonName  string        `json:"coljsonname"`
// 	DataType        string        `json:"datatype"`
// 	CharMaxLen      int           `json:"maxlen"`
// 	ColumnType      string        `json:"coltype"`
// 	Nump            int           `json:"nump"`
// 	Nums            int           `json:"nums"`
// 	Comment         string        `json:"comment"`
// 	DataTypeJava    string        `json:"datatypejava"`
// 	DataTypeGo      string        `json:"datatypego"`
// 	ColumnKey       string        `json:"columnkey"`
// 	Extra           string        `json:"extra"`
// 	OrdinalPosition string        `json:"position"`
// 	ModelTag        template.HTML `json:"modeltag"`
// 	ArgTag          template.HTML `json:"argtag"`
// }

// func (col *Column) IsKey() bool {
// 	return col.ColumnKey == "PRI"
// }

// func (col *Column) AutoIncrement() bool {
// 	return strings.Index(col.Extra, "auto_increment") > -1
// }

// func (col *Column) Build() string {
// 	return col.ColumnName + " " + col.DataType
// }

// type DstData struct {
// 	Package    string      `json:"package'`
// 	Model      string      `json:"model'`
// 	TableName  string      `json:"tablename"`
// 	ModelL     string      `json:"modell"`
// 	ModelApi   template.JS `json:"modelapi"`
// 	DefaultObj template.JS `json:"defaultobj"`
// 	Columns    []Column    `json:"columns"`
// 	Comment    string      `json:"comment"`
// 	Now        string      `json:"now"`
// }

// func ucfirst(str string) string {
// 	for i, v := range str {
// 		return string(unicode.ToUpper(v)) + str[i+1:]
// 	}
// 	return ""
// }
// func lcfirst(str string) string {
// 	for i, v := range str {
// 		return string(unicode.ToLower(v)) + str[i+1:]
// 	}
// 	return ""
// }

// //abc_def_ghi=> AbcDefGhi
// func transfer(in string) string {
// 	dstdata := make([]string, 0)
// 	inarr := strings.Split(in, "_")
// 	for _, v := range inarr {
// 		dstdata = append(dstdata, ucfirst(v))
// 	}
// 	return strings.Join(dstdata, "")
// }

// var datatypemapgo map[string]string = map[string]string{
// 	"tinyint":   "int",
// 	"smallint":  "int",
// 	"mediumint": "int",
// 	"int":       "int",
// 	"integer":   "int",
// 	"bigint":    "uint",
// 	"float":     "float64",
// 	"double":    "float64",

// 	"decimal":   "float64",
// 	"datetime":  "restgo.DateTime",
// 	"date":      "restgo.Date",
// 	"timestamp": "uint",
// 	"char":      "string",
// 	"varchar":   "string",
// 	"bit":       "bool",
// 	"numeric":   "float64",
// 	"text":      "string",
// }
