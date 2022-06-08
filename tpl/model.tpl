package model

type {{.VarName}}Model struct {
	*Model
}

func New{{.VarName}}Model() *{{.VarName}}Model {
	return &{{.VarName}}Model{
		NewModel("{{.TableName}}"),
	}
}
