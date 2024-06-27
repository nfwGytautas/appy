package appy_templates

const GoModTemplate = `
module {{.Domain}}/{{.ProjectName}}

replace github.com/nfwGytautas/appy-go => ../../../appy-go
`
