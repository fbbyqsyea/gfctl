package api

import (
	"{{.GetPackageName}}/contexts"
	"{{.GetPackageName}}/service"
	"github.com/gin-gonic/gin"
)

type {{.ApiName}}Api struct {
	*Api
	Svc *service.Service
}

func New{{.ApiName}}Api() *{{.ApiName}}Api {
	svc := service.New{{.ApiName}}Service()
	return &{{.ApiName}}Api{
		Api: NewApi(svc),
		Svc: svc,
	}
}
{{if .IsUseSwagger}}
// 新增{{.Comment}}信息
// @Summary 新增{{.Comment}}信息接口
// @Schemes http https
// @Description 用于新增{{.Comment}}信息
// @Tags {{.Comment}}相关
// @Produce json
// @param object body contexts.{{.ApiName}}InsertRequest true "{{.Comment}}信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "新增的{{.Comment}}id"
// @Router /api/{{.ApiName}}/insert [post]
{{end}}func (a *{{.ApiName}}Api) Insert(c *gin.Context) {
	a.Api.Insert(c, &contexts.{{.ApiName}}InsertRequest{})
}
{{if .IsUseSwagger}}
// 更新{{.Comment}}信息
// @Summary 更新{{.Comment}}信息接口
// @Schemes http https
// @Description 用于更新{{.Comment}}信息
// @Tags {{.Comment}}相关
// @Produce json
// @param object body contexts.{{.ApiName}}UpdateRequest true "{{.Comment}}信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "更新影响行数"
// @Router /api/{{.ApiName}}/update [post]
{{end}}func (a *{{.ApiName}}Api) Update(c *gin.Context) {
	a.Api.Update(c, &contexts.{{.ApiName}}UpdateRequest{})
}
{{if .IsUseSwagger}}
// 更新{{.Comment}}状态信息
// @Summary 更新{{.Comment}}状态信息接口
// @Schemes http https
// @Description 用于更新{{.Comment}}状态信息
// @Tags {{.Comment}}相关
// @Produce json
// @param object body contexts.StatusRequest true "状态信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "更新影响行数"
// @Router /api/{{.ApiName}}/status [post]
{{end}}func (a *{{.ApiName}}Api) Status(c *gin.Context) {
	a.Api.Status(c, &contexts.StatusRequest{})
}
{{if .IsUseSwagger}}
// 删除{{.Comment}}
// @Summary 删除{{.Comment}}接口
// @Schemes http https
// @Description 用于删除{{.Comment}}
// @Tags {{.Comment}}相关
// @Produce json
// @param object body contexts.RemoveRequest true "{{.Comment}}信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "删除影响行数"
// @Router /api/{{.ApiName}}/remove [post]
{{end}}func (a *{{.ApiName}}Api) Remove(c *gin.Context) {
	a.Api.Remove(c, &contexts.RemoveRequest{})
}
{{if .IsUseSwagger}}
// 批量删除{{.Comment}}
// @Summary 批量删除{{.Comment}}接口
// @Schemes http https
// @Description 用于批量删除{{.Comment}}
// @Tags {{.Comment}}相关
// @Produce json
// @param object body contexts.RemovesRequest true "{{.Comment}}信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "删除影响行数"
// @Router /api/{{.ApiName}}/removes [post]
{{end}}func (a *{{.ApiName}}Api) Removes(c *gin.Context) {
	a.Api.Removes(c, &contexts.RemovesRequest{})
}
{{if .IsUseSwagger}}
// 获取{{.Comment}}信息
// @Summary 获取{{.Comment}}信息接口
// @Schemes http https
// @Description 用于获取{{.Comment}}信息
// @Tags {{.Comment}}相关
// @Produce json
// @param id query uint true "{{.Comment}}ID"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=contexts.{{.ApiName}}Data}
// @Router /api/{{.ApiName}}/get [get]
{{end}}func (a *{{.ApiName}}Api) Get(c *gin.Context) {
	a.Api.Get(c, &contexts.GetRequest{}, &contexts.{{.ApiName}}Data{}, false)
}
{{if .IsUseSwagger}}
// 获取{{.Comment}}列表
// @Summary 获取{{.Comment}}列表接口
// @Schemes http https
// @Description 用于获取{{.Comment}}列表
// @Tags {{.Comment}}相关
// @Produce json
// @param id query uint false "{{.Comment}}id"
// @param status query int false "状态 1:启用 2:禁用 默认:1"
// @param page query uint64 false "页数 默认:1"
// @param limit query uint64 false "页数量 默认:20"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=[]contexts.{{.ApiName}}Data{}}
// @Router /api/{{.ApiName}}/list [get]
{{end}}
func (a *{{.ApiName}}Api) List(c *gin.Context) {
	a.Api.List(c, &contexts.{{.ApiName}}ListRequest{}, &[]contexts.{{.ApiName}}Data{}, false)
}