package appy_templates

const MainTemplate = `
package main

import (
	"github.com/nfwGytautas/appy-go"
	appy_http "github.com/nfwGytautas/appy-go/http"

	{{ .ProjectName }}_api "{{ .Domain }}/{{ .ProjectName }}/api"
)

{{ GeneratedNotice }}

func main() {
	options := appy.AppyOptions{
		Environment: appy.DefaultEnvironment(),
	}

	{{ if .Http }}
	// HTTP Options
	options.HTTP = &appy_http.HttpOptions{
		Address: "{{ .Http.Address }}",
		Mapper: nil,
		SSL: nil,
	}
	{{ end }}

	appy.Initialize(options)

	Initialize{{ .ProjectName }}()
	{{ .ProjectName }}_api.Setup()

	appy.Takeover()
}

`
