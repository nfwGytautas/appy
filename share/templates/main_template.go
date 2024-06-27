package appy_templates

const MainTemplate = `
package main

import (
	"github.com/nfwGytautas/appy-go"
	appy_http "github.com/nfwGytautas/appy-go/http"
	appy_jobs "github.com/nfwGytautas/appy-go/jobs"

	{{ .ProjectName }}_api "{{ .Domain }}/{{ .ProjectName }}/api"
	{{ .ProjectName }}_jobs "{{ .Domain }}/{{ .ProjectName }}/jobs"
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

	{{ if .Jobs }}
	// Job Options
	options.Jobs = &appy_jobs.JobSchedulerOptions{
		PoolTick: 1 * time.Second,
	}
	{{ end }}

	appy.Initialize(options)

	Initialize{{ .ProjectName }}()
	{{ if .Jobs }} {{ .ProjectName }}_jobs.Setup() {{ end }}
	{{ .ProjectName }}_api.Setup()

	appy.Takeover()
}

`
