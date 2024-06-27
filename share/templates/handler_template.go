package appy_templates

const HandlerTemplate = `
package {{ .Config.ProjectName }}_handlers

func {{ .Endpoint.Name | Title }}Handler(ctx context.Context, args {{ .Endpoint.Name | Title }}Args) {{ .Endpoint.Name | Title }}Result {
	// Your handler logic here
}

`

const HandlerImplementationTemplate = `
package {{ .ProjectName }}_handlers

{{ GeneratedNotice }}

{{ range .Http.Endpoints }}

type {{ .Name | Title }}Args struct {

}

type {{ .Name | Title }}Result struct {
	StatusCode int
	Body       interface{}
}

{{ end }}

`
