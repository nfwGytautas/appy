package appy_templates

const HttpWiringTemplate = `
package {{ .ProjectName }}_http

import (
	appy_http "github.com/nfwGytautas/appy-go/http"
)

{{ GeneratedNotice }}

func SetupHttpEndpoints() {
	// Setup groups
	root := appy_http.Get().Root()

	api := root.Group("/api")
	// ws := root.Group("/ws")

	// Endpoints
	{{ range .Http.Endpoints }}
	api.{{.Method}}("{{.Path}}", {{ .Name | Title }}Handler)
	{{ end }}
}

`
