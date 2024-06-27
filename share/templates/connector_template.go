package appy_templates

const ConnectorTemplate = `
package {{ .ProjectName}}_api

import (
	{{ .ProjectName }}_http "{{ .Domain }}/{{ .ProjectName }}/api/http"
)

{{ GeneratedNotice }}

func Setup() {
	{{ if .Http }}
	{{ .ProjectName }}_http.SetupHttpEndpoints()
	{{ end }}
}

`
