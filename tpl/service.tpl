package service

import (
	"{{.GetPackageName}}/model"
)

type {{.VarName}}Service struct {
	*Service
	Mdl *model.{{.VarName}}Model
}

func New{{.VarName}}Service() *{{.VarName}}Service {
	mdl := model.New{{.VarName}}Model()
	return &{{.VarName}}Service{
		Service: NewService(mdl),
		Mdl:     mdl,
	}
}
