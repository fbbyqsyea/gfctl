package router

import (
	"{{.GetPackageName}}/api"
	"{{.GetPackageName}}/middleware"
	"github.com/gin-gonic/gin"
)

type {{.VarName}}Router struct{}

func New{{.VarName}}Router() *{{.VarName}}Router {
	return &{{.VarName}}Router{}
}

func (r *{{.VarName}}Router) Register(g *gin.RouterGroup) {
	{{.Name}}g := g.Group("{{.Name}}")
	{
		{{.Name}}Api := api.New{{.VarName}}Api()
		{{if .IsUseAuth}}
		{{.Name}}g.Use(middleware.Authorization())
        {{end}}

		{{.Name}}g.POST("insert", {{.Name}}Api.Insert)
		{{.Name}}g.POST("update", {{.Name}}Api.Update)
		{{.Name}}g.POST("status", {{.Name}}Api.Status)
		{{.Name}}g.POST("remove", {{.Name}}Api.Remove)
		{{.Name}}g.POST("removes", {{.Name}}Api.Removes)
		{{.Name}}g.GET("get", {{.Name}}Api.Get)
		{{.Name}}g.GET("list", {{.Name}}Api.List)

	}
}